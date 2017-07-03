package http

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/multimfi/bot/pkg/irc/irctest"
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
		},
		{
			"endsAt": "0001-01-01T00:00:00Z",
			"generatorURL": "http://generator:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=0",
			"labels": {
				"alertname": "all_service_down",
				"instance": "127.0.0.1:9101",
				"job": "test"
			},
			"startsAt": "2017-05-26T18:30:00.001+03:00",
			"status": "firing"
		},
		{
			"endsAt": "0001-01-01T00:00:00Z",
			"generatorURL": "http://generator:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=0",
			"labels": {
				"alertname": "all_service_down",
				"instance": "127.0.0.1:9102",
				"job": "test"
			},
			"startsAt": "2017-05-26T18:30:00.002+03:00",
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
		},
		{
			"endsAt": "2017-05-27T18:30:00.000+03:00",
			"generatorURL": "http://generator:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=0",
			"labels": {
				"alertname": "all_service_down",
				"instance": "127.0.0.1:9101",
				"job": "test"
			},
			"startsAt": "2017-05-26T18:30:00.000+03:00",
			"status": "resolved"
		},
		{
			"endsAt": "2017-05-27T18:30:00.000+03:00",
			"generatorURL": "http://generator:9090/graph?g0.expr=up+%3D%3D+0\u0026g0.tab=0",
			"labels": {
				"alertname": "all_service_down",
				"instance": "127.0.0.1:9102",
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

var alertHash = [][16]uint8{
	[16]uint8{0x6d, 0x8f, 0x6d, 0xb0, 0x81, 0x1a, 0x36, 0x4c, 0x93, 0xd7, 0xe3, 0x35, 0x8f, 0x96, 0xff, 0xf0},
	[16]uint8{0xb2, 0x2a, 0x64, 0x48, 0xd2, 0x1, 0x1, 0x73, 0x53, 0x94, 0x64, 0xc2, 0xd3, 0xa4, 0xdb, 0xf0},
	[16]uint8{0xfc, 0xea, 0x44, 0xbf, 0xc0, 0xea, 0x13, 0xcb, 0x64, 0x25, 0x78, 0x9, 0x6b, 0x75, 0x39, 0x8f},
}

func newTestServer() (*Server, *httptest.Server) {
	m := http.NewServeMux()
	h := httptest.NewServer(m)
	i := irctest.NewClient("t", "t", "#t", "")

	ret := NewServer(i, m, nil)
	go ret.Dial()

	return ret, h
}

func alertBuf() *bytes.Buffer {
	return bytes.NewBuffer(testAlert)
}

func resolvBuf() *bytes.Buffer {
	return bytes.NewBuffer(testResolve)
}

func wait(t *testing.T, e string, f func() bool) {
	for x := 0; x < 5; x++ {
		if f() {
			return
		}
		time.Sleep(time.Millisecond * 10)
	}
	t.Fatal("wait:", e)
}

func TestAlertResolve(t *testing.T) {
	t.Parallel()
	srv, hts := newTestServer()
	url := hts.URL + "/alertmanager"
	defer hts.Close()

	if ret, err := http.DefaultClient.Post(url, "", alertBuf()); err != nil {
		t.Fatal("alert POST error", err, ret.Status)
	}

	wait(t, "alert not received", func() bool { return len(srv.pool.List()) == 3 })

	for k, v := range srv.pool.List() {
		if a := v.Hash(); a != alertHash[k] {
			t.Fatalf("hash mismatch %x != %x", a, alertHash[k])
		}
	}

	if ret, err := http.DefaultClient.Post(url, "", resolvBuf()); err != nil {
		t.Fatal("resolve POST error", err, ret)
	}

	wait(t, "alert not resolved", func() bool { return len(srv.pool.List()) == 0 })
}

func BenchmarkAlert(b *testing.B) {
	srv, hts := newTestServer()
	hts.Close()

	adata, err := unmarshal(alertBuf())
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		srv.alert(&adata.Alerts[0])
	}
}
