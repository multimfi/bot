package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	stdhttp "net/http"

	"github.com/multimfi/bot/http"
	"github.com/multimfi/bot/irc"
)

var buildversion = "devel"

type errFunc func() error

func fatal(fs ...errFunc) {
	for _, f := range fs {
		go func(f errFunc) {
			err := f()
			if err != nil {
				log.Fatal(err)
			}
		}(f)
	}
}

func config(file, tfile string) *http.Config {
	r := new(http.Config)

	f, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		log.Printf("config error: %v", err)
	} else if err != nil {
		log.Fatalf("config error: %v", err)
	} else {
		if err := json.Unmarshal(f, r); err != nil {
			log.Fatalf("config json error: %v", err)
		}
	}

	f, err = ioutil.ReadFile(tfile)
	if os.IsNotExist(err) {
		log.Printf("config template error: %v", err)
		return r
	}
	if err != nil {
		log.Fatalf("config template error: %v", err)
	}
	r.Template = string(f)
	return r
}

func version() string {
	return fmt.Sprintf("build: %s, runtime: %s", buildversion, runtime.Version())
}

var (
	flagConfig      = flag.String("cfg", "bot.json", "bot configuration file")
	flagTemplate    = flag.String("cfg.template", "template.tmpl", "template file")
	flagIRCServer   = flag.String("irc.server", "127.0.0.1:6667", "irc server address")
	flagIRCChannel  = flag.String("irc.channel", "#test", "irc channel to join")
	flagIRCUsername = flag.String("irc.user", "Bot", "irc username")
	flagIRCNickname = flag.String("irc.nick", "bot", "irc nickname")
	flagAMListen    = flag.String("alertmanager.addr", "127.0.0.1:9500", "alertmanager webhook listen address")
	flagVersion     = flag.Bool("version", false, "version")
)

func main() {
	log.SetFlags(log.Lshortfile)
	flag.Parse()

	if *flagVersion {
		fmt.Fprintln(os.Stderr, version())
		os.Exit(0)
	}

	mux := stdhttp.NewServeMux()

	hs := &stdhttp.Server{
		Addr:    *flagAMListen,
		Handler: mux,
	}
	ic := irc.NewClient(
		*flagIRCNickname,
		*flagIRCUsername,
		*flagIRCChannel,
		*flagIRCServer,
	)

	ic.Handle("!version", func(string) string {
		return version()
	})

	ic.Handle("!ping", func(string) string {
		return "pong"
	})

	srv := http.NewServer(ic, mux, config(*flagConfig, *flagTemplate))

	fatal(
		srv.Dial,
		hs.ListenAndServe,
	)

	ctx := context.Background()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-sig:
		log.Printf("received signal %s, shutting down", s)

		ic.Quit()

		ctx, cfunc := context.WithTimeout(ctx, time.Second*5)
		defer cfunc()
		if err := hs.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
	}
}
