package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"bitbucket.org/multimfi/bot/alert"
	"bitbucket.org/multimfi/bot/irc"
	"bitbucket.org/multimfi/bot/responder"
)

var ErrChanClosed = errors.New("alertmanager: closed channel")

type Server struct {
	irc     *irc.Client
	pool    *alert.Pool
	rpool   *responder.Pool
	spool   *subpool
	alertCh chan *alert.Alert
}

func NewServer(irc *irc.Client, srv *http.ServeMux) *Server {
	r := &Server{
		irc:     irc,
		pool:    alert.NewPool(),
		rpool:   responder.NewPool(),
		spool:   newPool(),
		alertCh: make(chan *alert.Alert, 2),
	}

	srv.HandleFunc("/alertmanager", r.alertManagerHandler)
	srv.HandleFunc("/ws", r.wsHandler)

	irc.StatusFunc(r.statusFunc)
	irc.Handle("!clear", r.clear)
	irc.Handle("!current", func(string) string {
		r := r.statusFunc()
		if r != "" {
			return r
		}
		return "no alerts"
	})

	irc.Handle("!reset", func(m string) string {
		r.rpool.ResetFailed(m)
		return "tada!"
	})

	irc.ActHandler(r.rpool.Update)

	srv.HandleFunc("/", r.statusPageHandler)

	return r
}

func (s *Server) clear(string) string {
	s.pool.Reset()
	return "tada!"
}

func (s *Server) alert(a *alert.Alert) {
	var (
		ca *alert.Alert
		m  string
		ok bool
	)

	ok, ca = s.pool.Add(a)
	oi := ca.Current()

	r := ca.Responders()

	n, i, err := s.rpool.Get(r)
	if err != nil {
		ok = true
		log.Printf("alert: error: %v", err)
	}

	if n != nil {
		ca.SetCurrent(int32(i))
		ok = ok || oi != int32(i)

		n.Ping()

		if n.Active() {
			m = ca.String()
		} else {
			m = fmt.Sprintf("%s - %s", ca, n.Name)
		}
	} else if len(r) < 1 {
		m = ca.String()
	} else {
		m = fmt.Sprintf("%s - %s", ca, ca.Responders())
	}

	if !ok {
		return
	}

	s.spool.broadcast(a)

	if err = s.irc.NSendMsg(m); err != nil {
		log.Printf("alert: error: %v", err)
	}
}

func (s *Server) resolve(a *alert.Alert) {
	if !s.pool.Remove(a) {
		return
	}

	s.spool.broadcast(a)

	if err := s.irc.NSendMsg(a.String()); err != nil {
		log.Printf("resolve: error: %v", err)
	}
}

func (s *Server) AlertManager() error {
	for a := range s.alertCh {
		switch a.Status {
		case alert.AlertFiring:
			s.alert(a)
		case alert.AlertResolved:
			s.resolve(a)
		default:
			log.Printf("alertmanager: invalid alert: %v", a)
		}
	}

	return ErrChanClosed
}

func unmarshal(data io.Reader) (*alert.Data, error) {
	r, err := ioutil.ReadAll(data)
	if err != nil {
		return nil, err
	}

	d := new(alert.Data)
	if err := json.Unmarshal(r, d); err != nil {
		return nil, err
	}

	return d, nil
}

func (s *Server) alertManagerHandler(w http.ResponseWriter, r *http.Request) {
	d, err := unmarshal(r.Body)
	if err != nil {
		log.Printf("handler: error: %v", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	r.Body.Close()

	for _, v := range d.Alerts {
		s.alertCh <- &v
	}
}

func (s *Server) statusPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(indexHTML)
}

func (s *Server) statusFunc() string {
	l := len(s.pool.List())
	if l > 0 {
		return fmt.Sprintf("alert: %d firing alerts", l)
	}
	return ""
}

func (s *Server) Dial() error {
	var d = time.Second

	for {
		err := s.irc.Dial()
		if err == irc.ErrDone {
			return nil
		}

		switch err.(type) {
		case irc.DialError:
			if d < time.Minute {
				d *= 2
			}
		default:
			d = time.Second
		}

		log.Printf("server: error %[1]q, %[1]T, reconnecting in %[2]s", err, d)
		time.Sleep(d)
	}
}
