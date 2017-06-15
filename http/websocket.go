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
	writeWait  = time.Second * 10
	pongWait   = time.Minute
	pingPeriod = (pongWait * 9) / 10
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var (
	jsTrue  = []byte(`{"status": true}`)
	jsFalse = []byte(`{"status": false}`)
)

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

type wsAlert struct {
	*alert.Alert
	Hash       uint32   `json:"hash"`
	Responders []string `json:"responders"`
	Current    int32    `json:"current"`
}

type wsResponder struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type wsTyped struct {
	Type string          `json:"type"`
	Msg  json.RawMessage `json:"msg"`
}

func wsAlertMsg(a *alert.Alert) ([]byte, error) {
	d := &wsAlert{
		Alert:      a,
		Hash:       a.Hash(),
		Responders: a.Responders(),
		Current:    a.Current(),
	}

	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	r := &wsTyped{
		Type: "alert",
		Msg:  j,
	}

	return json.Marshal(r)
}

func (s *Server) ircState() ([]byte, error) {
	r := &wsTyped{
		Type: "irc",
	}
	if s.irc.IsReady() {
		r.Msg = jsTrue
	} else {
		r.Msg = jsFalse
	}
	return json.Marshal(r)
}

func (s *Server) responderState() ([]byte, error) {
	lr := s.rpool.List()
	d := make([]*wsResponder, 0)

	for _, v := range lr {
		i := &wsResponder{Name: v.Name}
		if v.Active() {
			i.State = "active"
		} else if v.Failed() {
			i.State = "failed"
		} else {
			i.State = "unknown"
		}
		d = append(d, i)
	}

	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	r := &wsTyped{
		Type: "responders",
		Msg:  j,
	}

	return json.Marshal(r)
}

func (s *subpool) broadcast(b []byte) {
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

func (s *subpool) broadcastAlert(a *alert.Alert) {
	b, err := wsAlertMsg(a)
	if err != nil {
		log.Printf("websocket: alertjson error: %v", err)
		return
	}
	s.broadcast(b)
}

func (s *Server) broadcastIRC() {
	b, err := s.ircState()
	if err != nil {
		log.Printf("websocket: ircjson error: %v", err)
		return
	}
	s.spool.broadcast(b)
}

func (s *Server) broadcastResponders() {
	b, err := s.responderState()
	if err != nil {
		log.Printf("websocket: ircjson error: %v", err)
		return
	}
	s.spool.broadcast(b)
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
		case "alerts":
			for _, v := range s.pool.List() {
				r, err := wsAlertMsg(v)
				if err != nil {
					log.Printf("websocket: reader error %v", err)
					break
				}
				ac.ch <- r
			}
		case "irc":
			r, err := s.ircState()
			if err != nil {
				log.Printf("websocket: reader error %v", err)
				break
			}
			ac.ch <- r
		case "responders":
			r, err := s.responderState()
			if err != nil {
				log.Printf("websocket: reader error %v", err)
				break
			}
			ac.ch <- r

		default:
			log.Printf("invalid command %s from %q", string(msg), ws.RemoteAddr())
		}
	}
}
