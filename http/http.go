package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/multimfi/bot/alert"
	"github.com/multimfi/bot/irc"
	"github.com/multimfi/bot/responder"
)

// ReceiverGroup is a mapped group of receivers, nil is considered an empty group.
type ReceiverGroup map[string][]string

// Server tracks alerts received from prometheus alertmanager.
type Server struct {
	irc   irc.Client
	pool  *alert.Pool
	rpool *responder.Pool
	spool *subpool
	rg    ReceiverGroup
	tg    *telegram
	tmpl  *template.Template
}

// TData is passed to template.
type TData struct {
	A *alert.Alert
	R *responder.Responder
	H bool
}

// NewTData returns a new TData for template.
// tdata.H is set to false when alert has no configured responders.
// If responder r is nil tdata.R is a empty responder.
func NewTData(a *alert.Alert, r *responder.Responder) (tdata TData) {
	if r == nil {
		r = &responder.Responder{}
	}
	return TData{
		A: a,
		R: r,
		H: len(a.Responders) != 0,
	}
}

// NewServer registers and returns a new Server.
func NewServer(irc irc.Client, srv *http.ServeMux, cfg *Config) *Server {
	r := &Server{
		irc:   irc,
		pool:  alert.NewPool(),
		rpool: responder.NewPool(),
		spool: newPool(),
	}

	var err error
	t := DefaultTemplate

	if cfg != nil {
		if cfg.Template != "" {
			t = cfg.Template
		}

		if cfg.Telegram.BotID != "" && cfg.Telegram.ChatID != "" {
			r.tg = newTelegram(cfg.Telegram.BotID, cfg.Telegram.ChatID)
		}
		r.rg = cfg.Receivers
	}

	r.tmpl, err = template.New("").Parse(
		strings.Replace(t, "\n", " ", -1),
	)
	if err != nil {
		log.Fatalf("http: template error: %v", err)
	}

	srv.HandleFunc("/alertmanager", r.alertManagerHandler)
	srv.HandleFunc("/ws", r.wsHandler)
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
	ok, ca := s.pool.Add(a)
	if ca.AllFail() {
		return
	}

	n, i, err := s.rpool.Get(ca.Responders)
	if err != nil {
		ca.SetAllFail()
		log.Printf("alert: error: %v", err)
	}

	if n != nil {
		ok = ok || ca.Current() != i
		ca.SetCurrent(i)
	}
	if !ok {
		return
	}
	if n != nil {
		n.Ping(s.broadcastResponders)
	}

	// websocketpool broadcast
	s.spool.broadcastAlert(ca)

	buf := new(bytes.Buffer)
	var msg string

	if err := s.tmpl.Execute(buf, NewTData(ca, n)); err != nil {
		log.Printf("alert: template error: %v", err)
		msg = "error: " + err.Error()
	} else {
		msg = buf.String()
	}

	s.tg.broadcastTelegram(msg)

	if !s.irc.IsReady() {
		return
	}
	if err = s.irc.NSendMsg(msg); err != nil {
		log.Printf("alert: error: %v", err)
	}
}

func (s *Server) resolve(a *alert.Alert) {
	if !s.pool.Remove(a) {
		return
	}
	s.spool.broadcastAlert(a)

	n, _, err := s.rpool.Get(a.Responders)
	if err != nil {
		log.Printf("alert: error: %v", err)
	}

	buf := new(bytes.Buffer)
	var msg string

	if err := s.tmpl.Execute(buf, NewTData(a, n)); err != nil {
		log.Printf("resolve: template error: %v", err)
		msg = "error: " + err.Error()
	} else {
		msg = buf.String()
	}

	s.tg.broadcastTelegram(msg)

	if !s.irc.IsReady() {
		return
	}
	if err := s.irc.NSendMsg(msg); err != nil {
		log.Printf("resolve: error: %v", err)
	}
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

		switch v.Status {
		case alert.AlertFiring:
			s.alert(&p)
		case alert.AlertResolved:
			s.resolve(&p)
		default:
			log.Printf("alertmanager: invalid state: %s", v.Status)
		}
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
