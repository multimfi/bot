package irc

import (
	"errors"
	"log"
	"sync"

	"github.com/sorcix/irc"
)

var (
	ErrNotReady = errors.New("irc: client not ready")
	ErrOpen     = errors.New("irc: open connection")
	ErrDone     = errors.New("irc: quit")
)

type DialError string

func (d DialError) Error() string {
	return "irc: dial: " + string(d)
}

type HandlerError struct {
	msg string
	err string
}

func (d HandlerError) Error() string {
	return "irc: handler '" + d.msg + "': " + d.err
}

type MsgHandler func(string) string

type Client struct {
	conn    *irc.Conn
	nick    string
	user    string
	channel string
	server  string
	ready   chan struct{}
	done    chan struct{}

	statefunc func() string
	actfunc   func(string)

	msgMu  sync.RWMutex
	msgMap map[string]MsgHandler
}

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

func (c *Client) StateFunc(s func() string) {
	c.statefunc = s
}

func (c *Client) ActivityFunc(f func(string)) {
	c.actfunc = f
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

// TODO: context
func (c *Client) Send(cmd string, params []string) error {
	<-c.ready
	return c.send(cmd, params)
}

func (c *Client) NSend(cmd string, params []string) error {
	select {
	case <-c.ready:
		return c.send(cmd, params)
	default:
		return ErrNotReady
	}
}

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

func (c *Client) SendMsg(msg string) error {
	if msg == "" {
		return nil
	}
	return c.Send(irc.CmdPrivMsg, []string{c.channel, msg})
}

func (c *Client) NSendMsg(msg string) error {
	if msg == "" {
		return nil
	}
	return c.NSend(irc.CmdPrivMsg, []string{c.channel, msg})
}

func (c *Client) join() error {
	return c.send(irc.CmdJoin, []string{c.channel})
}

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

func (c *Client) IsReady() bool {
	select {
	case <-c.ready:
		return true
	default:
		return false
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
	c.actfunc(m.Name)
}

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

			if c.statefunc == nil {
				break
			}
			if err := c.SendMsg(c.statefunc()); err != nil {
				return err
			}

		// TODO: non-blocking?
		case irc.CmdPrivMsg:
			c.recActivity(m)

			if err := c.msgHandler(m); err != nil {
				log.Println(err)
			}
		}
	}
}
