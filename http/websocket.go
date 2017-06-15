package http

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"bitbucket.org/multimfi/bot/alert"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type subpool struct {
	mu sync.RWMutex
	id uint64
	m  map[uint64]*client
}

type client struct {
	id uint64
	ch chan []byte
}

func newPool() *subpool {
	return &subpool{
		m: make(map[uint64]*client),
	}
}

type wsData struct {
	*alert.Alert
	Hash       uint32   `json:"hash"`
	Responders []string `json:"responders"`
	Current    int32    `json:"current"`
}

func wsAlert(a *alert.Alert) ([]byte, error) {
	d := wsData{
		Alert:      a,
		Hash:       a.Hash(),
		Responders: a.Responders(),
		Current:    a.Current(),
	}
	return json.Marshal(&d)
}

func (s *subpool) broadcast(a *alert.Alert) {
	b, err := wsAlert(a)
	if err != nil {
		log.Printf("websocket: json error: %v", err)
		return
	}

	s.mu.RLock()
	for k, v := range s.m {
		select {
		case v.ch <- b:
		default:
			log.Printf("websocket: broadcast: %d failed", k)
		}
	}
	s.mu.RUnlock()
}

func (s *subpool) unsubscribe(c *client) {
	s.mu.Lock()
	if _, exists := s.m[c.id]; exists {
		close(c.ch)
		delete(s.m, c.id)
	}
	s.mu.Unlock()
}

func (s *subpool) subscribe() *client {
	s.mu.Lock()
	s.id++

	c := &client{
		id: s.id,
		ch: make(chan []byte, 4),
	}
	s.m[c.id] = c

	s.mu.Unlock()
	return c
}

func wsWriter(ac *client, ws *websocket.Conn) {
	pt := time.NewTicker(pingPeriod)
	defer pt.Stop()

	for {
		select {
		case a, ok := <-ac.ch:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := ws.WriteMessage(websocket.TextMessage, a); err != nil {
				return
			}

		case <-pt.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (s *Server) wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Printf("websocket: error: %v", err)
		}
		return
	}
	defer ws.Close()

	ws.SetReadLimit(32)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	ac := s.spool.subscribe()
	go wsWriter(ac, ws)
	defer s.spool.unsubscribe(ac)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}
		switch string(msg) {
		case "status":
			for _, v := range s.pool.List() {
				r, err := wsAlert(v)
				if err != nil {
					log.Printf("websocket: reader error %v", err)
				}
				select {
				case ac.ch <- r:
				default:
					log.Printf("websocket: client send blocked for %d, %q", ac.id, ws.RemoteAddr())
				}
			}
		default:
			log.Printf("invalid command %s from %q", string(msg), ws.RemoteAddr())
		}
	}
}
