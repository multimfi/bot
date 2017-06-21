package http

type Config struct {
	Receivers ReceiverGroup `json:"receivers"`
	Telegram  struct {
		BotID  string `json:"botid"`
		ChatID string `json:"chatid"`
	} `json:"telegram"`
}
