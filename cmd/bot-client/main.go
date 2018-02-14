package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	stdhttp "net/http"
	"os"
	"runtime"
	"strings"
	"text/template"

	"github.com/multimfi/bot/alert"
	"github.com/multimfi/bot/http"
)

var buildversion = "devel"

func handleAlert(d []byte) error {
	a := new(alert.Alert)
	if err := json.Unmarshal(d, a); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, http.NewTData(a, nil)); err != nil {
		return err
	}

	switch a.Status {
	case alert.AlertFiring:
		fmt.Println(a.Status, buf)

	case alert.AlertResolved:
		fmt.Println(a.Status, buf)

	default:
		return fmt.Errorf("invalid status %s", a.Status)
	}

	return nil
}

func loadTemplate(file string) *template.Template {
	var t string

	f, err := ioutil.ReadFile(file)
	if err == nil {
		t = string(f)
	} else if os.IsNotExist(err) {
		t = http.DefaultTemplate
		log.Printf("template: load error: %v", err)
	} else {
		log.Fatalf("template: load error: %v", err)
	}

	ret, err := template.New("").Parse(
		strings.Replace(t, "\n", " ", -1),
	)
	if err != nil {
		log.Fatalf("template: parse error: %v", err)
	}

	return ret
}

func version() string {
	return fmt.Sprintf("build: %s, runtime: %s", buildversion, runtime.Version())
}

var (
	flagAddress  = flag.String("sse.addr", "http://127.0.0.1:9500/sse", "sse endpoint")
	flagVersion  = flag.Bool("version", false, "print version")
	flagTemplate = flag.String("template", "template.tmpl", "template file")
)

var tmpl *template.Template

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	if *flagVersion {
		fmt.Fprintln(os.Stderr, version())
		os.Exit(0)
	}

	tmpl = loadTemplate(*flagTemplate)

	r, err := stdhttp.Get(*flagAddress)
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(r.Body)
	s.Split(splitFunc)

	for s.Scan() {
		e, err := parseSSE(s.Bytes())
		if err != nil {
			log.Println(err)
			continue
		}

		switch e.event {
		case http.EventAlert:
			err = handleAlert(e.data)
		case http.EventIRCDown:
			log.Println("irc down")
		case http.EventIRCReady:
			log.Println("irc ready")
		case http.EventResponder:
			log.Println("responder event")
		}

		if err != nil {
			log.Fatal(err)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}
