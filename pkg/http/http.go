package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/multimfi/bot/pkg/alert"
	"github.com/multimfi/bot/pkg/irc"
	"github.com/multimfi/bot/pkg/responder"
)

var errChanClosed = errors.New("alertmanager: closed channel")

// ReceiverGroup is a mapped group of receivers, nil is considered an empty group.
type ReceiverGroup map[string][]string

// Server tracks alerts received from prometheus alertmanager.
type Server struct {
	irc     irc.Client
	pool    *alert.Pool
	rpool   *responder.Pool
	spool   *subpool
	alertCh chan *alert.Alert
	rg      ReceiverGroup
	tg      *telegram
}

// NewServer registers and returns a new Server.
func NewServer(irc irc.Client, srv *http.ServeMux, cfg *Config) *Server {
	r := &Server{
		irc:     irc,
		pool:    alert.NewPool(),
		rpool:   responder.NewPool(),
		spool:   newPool(),
		alertCh: make(chan *alert.Alert, 2),
	}

	if cfg != nil {
		if cfg.Telegram.BotID != "" && cfg.Telegram.ChatID != "" {
			r.tg = newTelegram(cfg.Telegram.BotID, cfg.Telegram.ChatID)
		}
		r.rg = cfg.Receivers
	}

	srv.HandleFunc("/alertmanager", r.alertManagerHandler)
	srv.HandleFunc("/ws", r.wsHandler)
	srv.HandleFunc("/p", r.pollHandler)
	srv.HandleFunc("/", r.statusPageHandler)

	irc.Handle("!clear", r.clear)
	irc.Handle("!reset", r.reset)

	irc.StateFunc(r.statusFunc)
	irc.ActivityFunc(r.rpool.Update)

	return r
}

func (s *Server) reset(m string) string {
	if !s.rpool.ResetFailed(m) {
		return m + ": not in a failed state."
	}
	s.broadcastResponders()
	return m + ": tada!"
}

func (s *Server) clear(m string) string {
	if !s.pool.Reset() {
		return m + ": no alerts!"
	}
	return m + ": tada!"
}

func (s *Server) alert(a *alert.Alert) {
	var (
		ca *alert.Alert
		m  string
		ok bool
	)

	ok, ca = s.pool.Add(a)
	oi := ca.Current()

	r := ca.Responders

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
		m = fmt.Sprintf("%s - %s", ca, ca.Responders)
	}

	if !ok {
		return
	}

	if n != nil {
		n.Ping(s.broadcastResponders)
	}

	// websocketpool broadcast
	s.spool.broadcastAlert(ca)

	msg := "A " + m
	s.tg.broadcastTelegram(msg)

	if err = s.irc.NSendMsg(msg); err != nil {
		log.Printf("alert: error: %v", err)
	}
}

func (s *Server) resolve(a *alert.Alert) {
	if !s.pool.Remove(a) {
		return
	}

	s.spool.broadcastAlert(a)
	s.tg.broadcastTelegram("r " + a.String())

	if err := s.irc.NSendMsg("r " + a.String()); err != nil {
		log.Printf("resolve: error: %v", err)
	}
}

// AlertManager consumes alerts from alertCh.
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
	return errChanClosed
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

func (s *Server) rGet(n string) []string {
	if s.rg == nil {
		return nil
	}

	if r, exists := s.rg[n]; exists {
		return r
	}

	return nil
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
		v.Responders = s.rGet(d.Receiver)
		p := v
		s.alertCh <- &p
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

// Dial connects server to irc, reconnects on failure.
func (s *Server) Dial() error {
	var d = time.Second

	for {
		err := s.irc.Dial()
		if err == irc.ErrDone {
			s.broadcastIRC()
			return nil
		}

		switch err.(type) {
		case *net.OpError:
			if d < time.Minute*2 {
				d *= 2
			}
		default:
			d = time.Second
		}

		if d < time.Second*2 {
			s.broadcastIRC()
		}

		log.Printf("irc: error %q, reconnecting in %s", err, d)
		time.Sleep(d)
	}
}
