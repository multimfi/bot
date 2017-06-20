package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"bitbucket.org/multimfi/bot/pkg/alert"
	"bitbucket.org/multimfi/bot/pkg/http"

	"github.com/godbus/dbus"
	"github.com/gorilla/websocket"
)

var buildversion = "devel"

type client struct {
	dbus *dbus.Conn
	ws   *websocket.Conn
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
	var a alert.Alert
	if err := json.Unmarshal(d, &a); err != nil {
		return err
	}

	switch a.Status {
	case alert.AlertFiring:
		c.notify(a.Status, a.String(), 2, 0)
	case alert.AlertResolved:
		c.notify(a.Status, a.String(), 0, 15000)
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

func newClient(addr string) (*client, error) {
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
		dbus: dc,
		ws:   wc,
	}, nil
}

var (
	flagAddress = flag.String("ws.addr", "ws://127.0.0.1:9500/ws", "websocket address")
	flagVersion = flag.Bool("version", false, "print version")
)

func version() string {
	return fmt.Sprintf("build: %s, runtime: %s", buildversion, runtime.Version())
}

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	if *flagVersion {
		fmt.Fprintln(os.Stderr, version())
		os.Exit(0)
	}

	c, err := newClient(*flagAddress)
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
