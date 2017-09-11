package http

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/multimfi/bot/alert"

	"github.com/gorilla/websocket"
)

// Constants related to websocket communication.
const (
	WriteWait  = time.Second * 10
	PongWait   = time.Minute
	PingPeriod = (PongWait * 9) / 10
)

// Events communicated via websocket.
const (
	EventIRC = 1 << iota
	EventAlert
	EventResponder
)

// Possible states for responder.
const (
	StateActive = iota
	StateFailed
	StateUnknown
)

type wsAlert struct {
	*alert.Alert
	Hash    [16]byte `json:"h"`
	Current int      `json:"c"`
}

// WSTyped is a typed json message.
type WSTyped struct {
	Type uint8           `json:"t"`
	Msg  json.RawMessage `json:"m"`
}

func wsAlertMsg(a *alert.Alert) ([]byte, error) {
	d := &wsAlert{
		Alert:   a,
		Hash:    a.Hash(),
		Current: a.Current(),
	}

	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	r := &WSTyped{
		Type: EventAlert,
		Msg:  j,
	}

	return json.Marshal(r)
}

func (s *Server) ircState() ([]byte, error) {
	r := &WSTyped{
		Type: EventIRC,
		Msg:  []byte{'0'},
	}
	if s.irc.IsReady() {
		r.Msg[0] = '1'
	}
	return json.Marshal(r)
}

func (s *Server) responderState() ([]byte, error) {
	lr := s.rpool.List()
	if len(lr) < 1 {
		return nil, nil
	}
	d := make(map[string]uint8, 0)

	for _, v := range lr {
		if v.Active() {
			d[v.Name] = StateActive
		} else if v.Failed() {
			d[v.Name] = StateFailed
		} else {
			d[v.Name] = StateUnknown
		}
	}

	j, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}

	r := &WSTyped{
		Type: EventResponder,
		Msg:  j,
	}

	return json.Marshal(r)
}

// subpool is a pool of websocket subscriptions.
type subpool struct {
	mu  sync.RWMutex // guards id and seq
	seq uint64
	m   map[uint64]*client
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

func (s *subpool) unsubscribe(c *client) {
	s.mu.Lock()
	if _, exists := s.m[c.id]; exists {
		close(c.ch)
		delete(s.m, c.id)
	}
	s.mu.Unlock()
}

// subscribe returns a client with an unique id.
// uniqueness is guaranteed by always incrementing subpool sequence number.
func (s *subpool) subscribe() *client {
	s.mu.Lock()
	s.seq++

	c := &client{
		id: s.seq,
		ch: make(chan []byte, 4),
	}
	s.m[c.id] = c

	s.mu.Unlock()
	return c
}

// broadcast does a non-blocking send to all subscribed clients.
// does not guarantee that clients receive sent bytes.
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
	if b != nil {
		s.spool.broadcast(b)
	}
}

func wsWriter(ac *client, ws *websocket.Conn) {
	pt := time.NewTicker(PingPeriod)
	defer pt.Stop()

	for {
		select {
		case a, ok := <-ac.ch:
			ws.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := ws.WriteMessage(websocket.TextMessage, a); err != nil {
				return
			}

		case <-pt.C:
			ws.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

	// not expecting more than a single byte of data.
	ws.SetReadLimit(1)
	ws.SetReadDeadline(time.Now().Add(PongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(PongWait)); return nil })

	ac := s.spool.subscribe()
	go wsWriter(ac, ws)
	defer s.spool.unsubscribe(ac)

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		if len(msg) < 1 {
			log.Printf("early: invalid command %X from %q", msg, ws.RemoteAddr())
			continue
		}
		m := msg[0]

		if (m & EventIRC) != 0 {
			r, err := s.ircState()
			if err != nil {
				log.Printf("websocket: reader error %v", err)
				break
			}
			ac.ch <- r
		}

		if (m & EventAlert) != 0 {
			for _, v := range s.pool.List() {
				r, err := wsAlertMsg(v)
				if err != nil {
					log.Printf("websocket: reader error %v", err)
					break
				}
				ac.ch <- r
			}
		}

		if (m & EventResponder) != 0 {
			r, err := s.responderState()
			if err != nil {
				log.Printf("websocket: reader error %v", err)
				break
			}
			if r != nil {
				ac.ch <- r
			}
		}
	}
}

// pollHandler can be used as a simple "longpoll".
// TODO: bundled events, resumed poll.
func (s *Server) pollHandler(w http.ResponseWriter, r *http.Request) {
	ac := s.spool.subscribe()
	defer s.spool.unsubscribe(ac)

	var cc <-chan bool
	if c, ok := w.(http.CloseNotifier); ok {
		cc = c.CloseNotify()
	}

	w.(http.Flusher).Flush()

	for {
		select {
		case <-cc:
			return
		case a, ok := <-ac.ch:
			if !ok {
				return
			}
			w.Write(a)
			return
		}
	}
}
