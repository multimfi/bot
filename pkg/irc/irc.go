package irc

import (
	"errors"
	"log"
	"sync"

	"github.com/sorcix/irc"
)

var (
	// ErrNotReady is returned when current connection is not ready to receive packages.
	ErrNotReady = errors.New("irc: client not ready")

	// ErrDone is returned when client has quit.
	ErrDone = errors.New("irc: quit")
)

// DialError is returned when Dial fails connecting to server.
type DialError string

func (d DialError) Error() string {
	return "irc: dial: " + string(d)
}

// HandlerError is returned when a message handler fails a non-blocking send.
type HandlerError struct {
	msg string
	err string
}

func (d HandlerError) Error() string {
	return "irc: handler '" + d.msg + "': " + d.err
}

// MsgHandler takes a username, sends returned string as non-blocking.
type MsgHandler func(string) string

type ircClient struct {
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

// NewClient returns a new irc client.
func NewClient(nick, user, channel, server string) Client {
	return &ircClient{
		nick:    nick,
		user:    user,
		channel: channel,
		server:  server,
		ready:   make(chan struct{}),
		done:    make(chan struct{}),
		msgMap:  make(map[string]MsgHandler),
	}
}

// StateFunc registers a statefunction for client,
// function is ran when listen receives welcome msg.
func (c *ircClient) StateFunc(s func() string) {
	c.statefunc = s
}

// ActivityFunc tracks activity for user.
func (c *ircClient) ActivityFunc(f func(string)) {
	c.actfunc = f
}

// Handle registers the handler for the given string.
func (c *ircClient) Handle(msg string, f MsgHandler) {
	c.msgMu.Lock()
	c.msgMap[msg] = f
	c.msgMu.Unlock()
}

func (c *ircClient) send(cmd string, params []string) error {
	msg := &irc.Message{
		Command: cmd,
		Params:  params,
	}
	return c.conn.Encode(msg)
}

// Send waits for connection readiness before sending.
func (c *ircClient) Send(cmd string, params []string) error {
	<-c.ready
	return c.send(cmd, params)
}

// NSend is a non-blocking version of Send,
// returns ErrNotReady if connection is not ready.
func (c *ircClient) NSend(cmd string, params []string) error {
	select {
	case <-c.ready:
		return c.send(cmd, params)
	default:
		return ErrNotReady
	}
}

func (c *ircClient) register() error {
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

// Dial connects and registers to server, blocks until error.
func (c *ircClient) Dial() error {
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

func (c *ircClient) pongHandler(msg *irc.Message) error {
	return c.send(irc.CmdPong, msg.Params)
}

// SendMsg sends msg to channel in client config.
func (c *ircClient) SendMsg(msg string) error {
	if msg == "" {
		return nil
	}
	return c.Send(irc.CmdPrivMsg, []string{c.channel, msg})
}

// NSendMsg is a nonblocking version of SendMsg.
func (c *ircClient) NSendMsg(msg string) error {
	if msg == "" {
		return nil
	}
	return c.NSend(irc.CmdPrivMsg, []string{c.channel, msg})
}

func (c *ircClient) join() error {
	return c.send(irc.CmdJoin, []string{c.channel})
}

// Quit closes the current connection.
func (c *ircClient) Quit() error {
	defer c.quit()

	// blocks requests until reconnected
	c.ready = make(chan struct{})

	if c.conn != nil {
		return c.send(irc.CmdQuit, []string{"shutdown"})
	}

	return nil
}

func (c *ircClient) quit() {
	select {
	case <-c.done:
		break
	default:
		close(c.done)
	}
}

func (c *ircClient) setready() {
	select {
	case <-c.ready:
		break
	default:
		close(c.ready)
	}
}

// IsReady returns the connection readiness.
func (c *ircClient) IsReady() bool {
	select {
	case <-c.ready:
		return true
	default:
		return false
	}
}

func (c *ircClient) msgHandler(m *irc.Message) error {
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

func (c *ircClient) recActivity(m *irc.Message) {
	c.actfunc(m.Name)
}

// TODO: context
func (c *ircClient) listen() error {
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
