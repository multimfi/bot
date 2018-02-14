package http

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/multimfi/bot/alert"
	"github.com/multimfi/bot/event"
)

const (
	EventIRCReady = iota
	EventIRCDown
	EventReset
	EventAlert
	EventResponder
)

const (
	StateActive = iota
	StateFailed
	StateUnknown
)

func (s *Server) ircstate() uint32 {
	if s.irc.IsReady() {
		return EventIRCReady
	} else {
		return EventIRCDown
	}
}

func (s *Server) broadcastIRC() {
	s.epool.Add(nil, s.ircstate())
}

func (s *Server) responderstate() []byte {
	lr := s.rpool.List()
	if len(lr) < 1 {
		return nil
	}

	d := make(map[string]uint8, len(lr))

	for _, v := range lr {
		if v.Active() {
			d[v.Name] = StateActive
		} else if v.Failed() {
			d[v.Name] = StateFailed
		} else {
			d[v.Name] = StateUnknown
		}
	}

	r, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	return r
}

func (s *Server) broadcastResponders() {
	if r := s.responderstate(); r != nil {
		s.epool.Add(r, EventResponder)
	}
}

type alertmsg struct {
	*alert.Alert
	Hash    string `json:"h"`
	Current int    `json:"c"`
}

func alertjson(a *alert.Alert) ([]byte, error) {
	hash := a.Hash()
	return json.Marshal(&alertmsg{
		Alert:   a,
		Hash:    hex.EncodeToString(hash[:]),
		Current: a.Current(),
	})
}

func (s *Server) broadcastAlert(a *alert.Alert) {
	b, err := alertjson(a)
	if err != nil {
		panic(err)
	}
	s.epool.Add(b, EventAlert)
}

var msgReset = []byte("id:0\nevent:2\ndata:\n\n")

func sse(w http.ResponseWriter, id uint32, event uint32, data []byte) error {
	_, err := fmt.Fprintf(w, "id:%d\nevent:%d\ndata:%s\n\n", id, event, data)
	return err
}

func resetjs(w http.ResponseWriter) error {
	_, err := w.Write(msgReset)
	w.(http.Flusher).Flush()
	return err
}

func (s *Server) current(w http.ResponseWriter) (gen uint32, err error) {
	if err := resetjs(w); err != nil {
		return 0, err
	}

	if err := sse(w, 0, s.ircstate(), nil); err != nil {
		return 0, err
	}

	if r := s.responderstate(); r != nil {
		if err := sse(w, 0, EventResponder, r); err != nil {
			return 0, err
		}
	}

	gen = s.epool.Lock()
	defer s.epool.Unlock()

	for _, v := range s.pool.List() {
		r, err := alertjson(v)
		if err != nil {
			return 0, err
		}

		if err := sse(w, 0, EventAlert, r); err != nil {
			return 0, err
		}
	}

	w.(http.Flusher).Flush()
	return gen, nil
}

func (s *Server) sserange(w http.ResponseWriter, last, cur uint32) error {
	for {
		if last == cur {
			break
		}

		last++

		r := s.epool.Get(last)
		if err := sse(w, r.Gen, r.Type, r.Data); err != nil {
			return err
		}
	}

	w.(http.Flusher).Flush()
	return nil
}

func sub(a, b uint32) uint32 {
	const max = ^uint32(0)

	if a > b {
		return (a - max) + b
	}

	return b - a
}

func (s *Server) ssehandler(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()

	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	hdr := w.Header()
	hdr.Set("Content-Type", "text/event-stream")
	hdr.Set("Cache-Control", "no-cache")
	hdr.Set("Connection", "keep-alive")

	var (
		last, gen uint32
		err       error
		ok        bool
	)

	last, err = s.current(w)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cc := r.Context().Done()
	go func() {
		<-cc
		s.epool.Broadcast()
	}()

	for {
		gen, ok = s.epool.Next(last, cc)
		if !ok {
			break
		}

		if sub(last, gen) > event.Len {
			last, err = s.current(w)
			break
		}

		s.sserange(w, last, gen)
		last = gen
	}

	if err != nil {
		log.Println("sse:", err)
	}
}
