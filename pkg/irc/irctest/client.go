package irctest

import "github.com/multimfi/bot/pkg/irc"

type client struct {
}

func NewClient(nick, user, channel, server string) *client {
	return &client{}
}

func (c *client) Dial() error   { select {} }
func (c *client) Quit() error   { return nil }
func (c *client) IsReady() bool { return true }

func (c *client) NSend(cmd string, params []string) error { return nil }
func (c *client) NSendMsg(msg string) error               { return nil }

func (c *client) Send(cmd string, params []string) error { return nil }
func (c *client) SendMsg(msg string) error               { return nil }

func (c *client) StateFunc(s func() string)           {}
func (c *client) ActivityFunc(f func(string))         {}
func (c *client) Handle(msg string, f irc.MsgHandler) {}
