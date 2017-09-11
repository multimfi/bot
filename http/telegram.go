package http

import (
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const (
	tapiTimeout = time.Second * 10
	tapiURL     = "https://api.telegram.org/"
)

type telegram struct {
	c      *http.Client
	apiURL string
	cID    string

	once sync.Once
	msg  chan string
	done chan struct{}
}

func newTelegram(bID string, cID string) *telegram {
	return &telegram{
		cID:    cID,
		apiURL: tapiURL + bID + "/sendMessage",
		msg:    make(chan string, 4),
		done:   make(chan struct{}),
		c: &http.Client{
			Timeout: tapiTimeout,
		},
	}
}

func (t *telegram) msgFmt(msg string) url.Values {
	r := make(url.Values)
	r.Add("chat_id", t.cID)
	r.Add("text", msg)
	return r
}

func (t *telegram) sendMessage(msg string) {
	r, err := t.c.PostForm(t.apiURL, t.msgFmt(msg))
	if err != nil {
		log.Printf("telegram: send: error %v", err)
		return
	}
	r.Body.Close()
	if r.StatusCode != http.StatusOK {
		log.Printf("telegram: send: error %d: %s", r.StatusCode, r.Status)
	}
}

func (t *telegram) sender() {
	for m := range t.msg {
		t.sendMessage(m)
	}

	log.Println("telegram: sender: channel closed")
	close(t.done)
}

func (t *telegram) broadcastTelegram(a string) {
	if t == nil {
		return
	}

	t.once.Do(func() { go t.sender() })

	select {
	case <-t.done:
		return
	case t.msg <- a:
	default:
		log.Println("telegram: channel full")
		<-t.msg
	}
}
