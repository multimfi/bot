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

	"bitbucket.org/multimfi/bot/pkg/alert"
	"bitbucket.org/multimfi/bot/pkg/irc"
	"bitbucket.org/multimfi/bot/pkg/responder"
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
	srv.HandleFunc("/p", r.pollHandler)

	irc.StateFunc(r.statusFunc)
	irc.Handle("!clear", r.clear)
	irc.Handle("!reset", func(m string) string {
		if !r.rpool.ResetFailed(m) {
			return m + ": not in a failed state."
		}
		r.broadcastResponders()
		return m + ": tada!"
	})

	irc.ActivityFunc(r.rpool.Update)

	srv.HandleFunc("/", r.statusPageHandler)

	return r
}

func (s *Server) clear(m string) string {
	if s.pool.Reset() {
		return m + ": tada!"
	}
	return m + ": no alerts!"
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
		log.Printf("alert: error: %v", err)
	}

	if n != nil {
		ca.SetCurrent(int32(i))
		ok = ok || oi != int32(i)

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

	if n != nil {
		n.Ping(s.broadcastResponders)
	}

	s.spool.broadcastAlert(ca)

	if err = s.irc.NSendMsg(m); err != nil {
		log.Printf("alert: error: %v", err)
	}
}

func (s *Server) resolve(a *alert.Alert) {
	if !s.pool.Remove(a) {
		return
	}

	s.spool.broadcastAlert(a)

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
	s.broadcastIRC()
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
			s.broadcastIRC()
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

		if d < time.Second*2 {
			s.broadcastIRC()
		}

		log.Printf("server: error %[1]q, %[1]T, reconnecting in %[2]s", err, d)
		time.Sleep(d)
	}
}
