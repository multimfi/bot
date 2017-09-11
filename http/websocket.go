package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/multimfi/bot/alert"

	"github.com/gorilla/websocket"
)

var errMaxClients = errors.New("subpool: allocate error: maxclients")

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

const (
	clientMax     = 256
	clientBufSize = 4
	poolBufSize   = 32
)

// subpool is a pool of websocket subscriptions.
type subpool struct {
	ch chan []byte

	mu  sync.Mutex
	seq uint32
	cc  [clientMax]chan []byte
}

func newPool() *subpool {
	r := &subpool{
		ch: make(chan []byte, poolBufSize),
	}
	go func() {
		for {
			a, ok := <-r.ch
			if !ok {
				return
			}

			r.mu.Lock()
			for n, c := range r.cc {
				if c == nil {
					continue
				}

				select {
				case c <- a:
					continue
				default:
				}

				log.Printf("subpool: client %d, channel full", n)
			}
			r.mu.Unlock()
		}
	}()
	return r
}

func get(w, l uint32) uint32 {
	if w <= l {
		return w - 1
	}

	w = w % l
	if w < 1 {
		w = l
	}
	return w - 1
}

func (s *subpool) allocate() (chan []byte, uint32, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seq++
	n := get(s.seq, clientMax)

	if s.cc[n] == nil {
		s.cc[n] = make(chan []byte, clientBufSize)
		return s.cc[n], n, nil
	}

	for i, c := range s.cc {
		if c != nil {
			continue
		}
		s.cc[i] = make(chan []byte, clientBufSize)
		return s.cc[i], uint32(i), nil
	}

	return nil, 0, errMaxClients
}

func (s *subpool) free(n uint32, ch chan []byte) {
	s.mu.Lock()
	s.cc[n] = nil
	s.mu.Unlock()
	close(ch)
}

// broadcast does a non-blocking send to all subscribed clients.
// does not guarantee that clients receive sent bytes.
func (s *subpool) broadcast(b []byte) {
	select {
	case s.ch <- b:
	default:
		log.Println("websocket: broadcast: channel full")
	}
}

func (s *subpool) broadcastAlert(a *alert.Alert) {
	b, err := wsAlertMsg(a)
	if err != nil {
		log.Printf("broadcast: alert: error: %v", err)
		return
	}
	s.broadcast(b)
}

func (s *Server) broadcastIRC() {
	b, err := s.ircState()
	if err != nil {
		log.Printf("broadcast: irc: error: %v", err)
		return
	}
	s.spool.broadcast(b)
}

func (s *Server) broadcastResponders() {
	b, err := s.responderState()
	if err != nil {
		log.Printf("broadcast: responders: error: %v", err)
		return
	}
	if b != nil {
		s.spool.broadcast(b)
	}
}

func wsWriter(c <-chan []byte, ws *websocket.Conn) {
	pt := time.NewTicker(PingPeriod)
	defer pt.Stop()

	for {
		select {
		case b, ok := <-c:
			ws.SetWriteDeadline(time.Now().Add(WriteWait))
			if !ok {
				ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			if err := ws.WriteMessage(websocket.TextMessage, b); err != nil {
				log.Printf("websocket: client error: %v", err)
				return
			}
		case <-pt.C:
			ws.SetWriteDeadline(time.Now().Add(WriteWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("websocket: client error: %v", err)
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
	ch, n, err := s.spool.allocate()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer s.spool.free(n, ch)

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
	go wsWriter(ch, ws)

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
				log.Printf("reader error %v", err)
				break
			}
			ch <- r
		}

		if (m & EventAlert) != 0 {
			for _, v := range s.pool.List() {
				r, err := wsAlertMsg(v)
				if err != nil {
					log.Printf("reader error %v", err)
					break
				}
				ch <- r
			}
		}

		if (m & EventResponder) != 0 {
			r, err := s.responderState()
			if err != nil {
				log.Printf("reader error %v", err)
				break
			}
			if r != nil {
				ch <- r
			}
		}
	}
}
