package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"bitbucket.org/multimfi/bot/pkg/irc/irctest"
)

var testAlert = []byte(`
{
	"alerts": [
		{
			"endsAt": "0001-01-01T00:00:00Z",
			"generatorURL": "http://generator:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=0",
			"labels": {
				"alertname": "all_service_down",
				"instance": "127.0.0.1:9100",
				"job": "test"
			},
			"startsAt": "2017-05-26T18:30:00.000+03:00",
			"status": "firing"
		}
	],
	"receiver": "group1",
	"status": "firing"
}
`)
var testResolve = []byte(`
{
	"alerts": [
		{
			"endsAt": "2017-05-27T18:30:00.000+03:00",
			"generatorURL": "http://generator:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=0",
			"labels": {
				"alertname": "all_service_down",
				"instance": "127.0.0.1:9100",
				"job": "test"
			},
			"startsAt": "2017-05-26T18:30:00.000+03:00",
			"status": "resolved"
		}
	],
	"receiver": "group1",
	"status": "resolved"
}
`)

const alertHash = 2194928353

func newTestServer() (*Server, *httptest.Server) {
	m := http.NewServeMux()
	h := httptest.NewServer(m)
	i := irctest.NewClient("t", "t", "#t", "")

	ret := NewServer(i, m, nil)
	go ret.Dial()
	go ret.AlertManager()

	return ret, h
}

func alertBuf() *bytes.Buffer {
	return bytes.NewBuffer(testAlert)
}

func resolvBuf() *bytes.Buffer {
	return bytes.NewBuffer(testResolve)
}

func TestAlertResolve(t *testing.T) {
	t.Parallel()
	srv, hts := newTestServer()
	url := hts.URL + "/alertmanager"

	if ret, err := http.DefaultClient.Post(url, "", alertBuf()); err != nil {
		t.Fatal("alert POST error", err, ret.Status)
	}

	p := srv.pool.List()
	if len(p) < 1 {
		t.Fatal("alert not received")
	}
	if a := p[0].Hash(); a != alertHash {
		t.Fatalf("hash mismatch %d != %d", a, alertHash)
	}

	if ret, err := http.DefaultClient.Post(url, "", resolvBuf()); err != nil {
		t.Fatal("resolve POST error", err, ret)
	}

	if len(srv.pool.List()) > 0 {
		t.Fatal("alert not resolved")
	}
}
