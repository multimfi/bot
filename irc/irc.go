package irc

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/sorcix/irc"
)

// Errors XXX
var (
	ErrNotReady = errors.New("irc: client not ready")
	ErrOpen     = errors.New("irc: open connection")
	ErrDone     = errors.New("irc: quit")
)

type (
	MsgHandler func(string) string
	StatusFunc func() string
)

// DialError XXX
type DialError string

func (d DialError) Error() string {
	return "irc: dial: " + string(d)
}

// HandlerError XXX
type HandlerError struct {
	msg string
	err string
}

func (d HandlerError) Error() string {
	return "irc: handler '" + d.msg + "': " + d.err
}

func truncate(t time.Duration) time.Duration {
	return t - t%time.Second
}

// Client XXX
type Client struct {
	conn     *irc.Conn
	nick     string
	user     string
	channel  string
	server   string
	statusfn StatusFunc
	ready    chan struct{}
	done     chan struct{}

	msgMu  sync.RWMutex
	msgMap map[string]MsgHandler

	actHandler func(string)
}

// NewClient XXX
func NewClient(nick, user, channel, server string) *Client {
	return &Client{
		nick:    nick,
		user:    user,
		channel: channel,
		server:  server,
		ready:   make(chan struct{}),
		done:    make(chan struct{}),
		msgMap:  make(map[string]MsgHandler),
	}
}

// StatusFunc XXX
func (c *Client) StatusFunc(s StatusFunc) {
	c.statusfn = s
}

func (c *Client) ActHandler(f func(string)) {
	c.actHandler = f
}

// Handle registers the handler for the given string.
func (c *Client) Handle(msg string, f MsgHandler) {
	c.msgMu.Lock()
	c.msgMap[msg] = f
	c.msgMu.Unlock()
}

func (c *Client) send(cmd string, params []string) error {
	msg := &irc.Message{
		Command: cmd,
		Params:  params,
	}
	return c.conn.Encode(msg)
}

// Send XXX
// TODO: context
func (c *Client) Send(cmd string, params []string) error {
	<-c.ready
	return c.send(cmd, params)
}

// NSend XXX
func (c *Client) NSend(cmd string, params []string) error {
	select {
	case <-c.ready:
		return c.send(cmd, params)
	default:
		return ErrNotReady
	}
}

// Register XXX
func (c *Client) register() error {
	if err := c.send(irc.CmdUser, []string{
		c.user, "0", "*", ":" + c.user,
	}); err != nil {
		return err
	}

	if err := c.send(irc.CmdNick, []string{c.nick}); err != nil {
		return err
	}

	return nil
}

// Dial XXX
func (c *Client) Dial() error {
	if c.conn != nil {
		c.conn.Close()
	}

	select {
	case <-c.done:
		return ErrDone
	default:
	}

	ic, err := irc.Dial(c.server)
	if err != nil {
		return DialError(err.Error())
	}
	c.conn = ic

	if err := c.listen(); err != nil {
		c.ready = make(chan struct{})
		c.conn.Close()
		return err
	}

	return nil
}

func (c *Client) pongHandler(msg *irc.Message) error {
	return c.send(irc.CmdPong, msg.Params)
}

// SendMsg XXX
func (c *Client) SendMsg(msg string) error {
	if msg == "" {
		return nil
	}
	return c.Send(irc.CmdPrivMsg, []string{c.channel, msg})
}

// NSendMsg XXX
func (c *Client) NSendMsg(msg string) error {
	if msg == "" {
		return nil
	}
	return c.NSend(irc.CmdPrivMsg, []string{c.channel, msg})
}

// Join XXX
func (c *Client) join() error {
	return c.send(irc.CmdJoin, []string{c.channel})
}

// Quit XXX
func (c *Client) Quit() error {
	defer c.quit()

	// blocks requests until reconnected
	c.ready = make(chan struct{})

	if c.conn != nil {
		return c.send(irc.CmdQuit, []string{"shutdown"})
	}

	return nil
}

func (c *Client) quit() {
	select {
	case <-c.done:
		break
	default:
		close(c.done)
	}
}

func (c *Client) setready() {
	select {
	case <-c.ready:
		break
	default:
		close(c.ready)
	}
}

func (c *Client) msgHandler(m *irc.Message) error {
	var (
		dst string
		msg string
		ret []string
	)

	dst, msg = m.Params[0], m.Params[1]

	h, ok := c.msgMap[msg]
	if !ok {
		return nil
	}

	if dst[0] != '#' {
		dst = m.Name
	}

	ret = []string{dst, h(m.Name)}
	if err := c.NSend(irc.CmdPrivMsg, ret); err != nil {
		return HandlerError{msg, err.Error()}
	}

	return nil
}

func (c *Client) recActivity(m *irc.Message) {
	c.actHandler(m.Name)
}

// Listen XXX
// TODO: context
func (c *Client) listen() error {
	if err := c.register(); err != nil {
		return err
	}

	for {
		m, err := c.conn.Decode()
		if err != nil {
			return err
		}

		select {
		case <-c.done:
			return ErrDone
		default:
		}

		switch m.Command {
		case irc.CmdPing:
			if err := c.pongHandler(m); err != nil {
				return err
			}

		case irc.ReplyWelcome:
			c.setready()
			if err := c.join(); err != nil {
				return err
			}

			if c.statusfn == nil {
				break
			}
			if err := c.SendMsg(c.statusfn()); err != nil {
				return err
			}

		// TODO: non-blocking?
		case irc.CmdPrivMsg:
			c.recActivity(m)

			if err := c.msgHandler(m); err != nil {
				log.Println(err)
			}
		}

		log.Printf("irc: server: %v: %v", m.Command, m)
	}
}
