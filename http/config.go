package http

type Config struct {
	Receivers ReceiverGroup `json:"receivers"`
	Template  string        `json:"-"`
}
