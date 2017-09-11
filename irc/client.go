package irc

// Client XXX.
type Client interface {
	Dial() error
	Quit() error

	NSend(cmd string, params []string) error
	NSendMsg(msg string) error

	Send(cmd string, params []string) error
	SendMsg(msg string) error

	StateFunc(s func() string)
	ActivityFunc(f func(string))

	IsReady() bool
	Handle(msg string, f MsgHandler)
}
