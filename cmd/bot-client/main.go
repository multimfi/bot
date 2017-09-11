package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"text/template"
	"time"

	"github.com/multimfi/bot/alert"
	"github.com/multimfi/bot/http"

	"github.com/godbus/dbus"
	"github.com/gorilla/websocket"
)

var buildversion = "devel"

type client struct {
	dbus     *dbus.Conn
	ws       *websocket.Conn
	template *template.Template
}

func (c *client) notify(title string, message string, urgency byte, timeout int32) {
	obj := c.dbus.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	call := obj.Call(
		"org.freedesktop.Notifications.Notify",
		0,
		"",
		uint32(0),
		"",
		title,
		message,
		[]string{},
		map[string]dbus.Variant{"urgency": dbus.MakeVariant(urgency)},
		timeout,
	)

	if call.Err != nil {
		log.Fatal(call.Err)
	}
}

func (c *client) handleAlert(d []byte) error {
	a := new(alert.Alert)
	if err := json.Unmarshal(d, a); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	var msg string

	err := c.template.Execute(buf, http.NewTData(a, nil))
	if err != nil {
		msg = "error: " + err.Error()
	} else {
		msg = buf.String()
	}

	switch a.Status {
	case alert.AlertFiring:
		c.notify(a.Status, msg, 2, 0)
	case alert.AlertResolved:
		c.notify(a.Status, msg, 0, 15000)
	}

	return nil
}

func (c *client) msgHandler(d []byte) error {
	var t http.WSTyped
	if err := json.Unmarshal(d, &t); err != nil {
		return err
	}

	var err error

	switch t.Type {
	case http.EventAlert:
		err = c.handleAlert(t.Msg)
	case http.EventIRC:
	case http.EventResponder:
		// TODO
	}

	return err
}

func loadTemplate(file string) *template.Template {
	var t string

	f, err := ioutil.ReadFile(file)
	if err == nil {
		t = string(f)
	} else if os.IsNotExist(err) {
		t = http.DefaultTemplate
		log.Printf("template: load error: %v", err)
	} else {
		log.Fatalf("template: load error: %v", err)
	}

	ret, err := template.New("").Parse(
		strings.Replace(t, "\n", " ", -1),
	)
	if err != nil {
		log.Fatalf("template: parse error: %v", err)
	}

	return ret
}

func newClient(addr, tmpl string) (*client, error) {
	dc, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}

	d := &websocket.Dialer{}
	wc, _, err := d.Dial(addr, nil)
	if err != nil {
		return nil, err
	}

	return &client{
		dbus:     dc,
		ws:       wc,
		template: loadTemplate(tmpl),
	}, nil
}

func version() string {
	return fmt.Sprintf("build: %s, runtime: %s", buildversion, runtime.Version())
}

var (
	flagAddress  = flag.String("ws.addr", "ws://127.0.0.1:9500/ws", "websocket address")
	flagVersion  = flag.Bool("version", false, "print version")
	flagTemplate = flag.String("template", "template.tmpl", "template file")
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	if *flagVersion {
		fmt.Fprintln(os.Stderr, version())
		os.Exit(0)
	}

	c, err := newClient(*flagAddress, *flagTemplate)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.ws.WriteMessage(websocket.TextMessage, []byte{http.EventAlert}); err != nil {
		log.Fatal(err)
	}

	c.ws.SetPingHandler(func(m string) error {
		err := c.ws.SetReadDeadline(time.Now().Add(http.PongWait))
		if err != nil {
			return err
		}
		err = c.ws.WriteControl(websocket.PongMessage, []byte(m), time.Now().Add(http.WriteWait))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})

	for {
		_, p, err := c.ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		if err := c.msgHandler(p); err != nil {
			log.Fatal(err)
		}
	}
}
