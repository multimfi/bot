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

	"bitbucket.org/multimfi/bot/pkg/http"
	"bitbucket.org/multimfi/bot/pkg/irc"
)

var buildversion = "devel"

var (
	flagConfig      = flag.String("cfg", "bot.json", "bot configuration file")
	flagIRCServer   = flag.String("irc.server", "127.0.0.1:6667", "irc server address")
	flagIRCChannel  = flag.String("irc.channel", "#test", "irc channel to join")
	flagIRCUsername = flag.String("irc.user", "Bot", "irc username")
	flagIRCNickname = flag.String("irc.nick", "bot", "irc nickname")
	flagAMListen    = flag.String("alertmanager.addr", "127.0.0.1:9500", "alertmanager webhook listen address")
	flagVersion     = flag.Bool("version", false, "version")
)

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

func version() string {
	return fmt.Sprintf("build: %s, runtime: %s", buildversion, runtime.Version())
}

func config(file string) *http.Config {
	r := new(http.Config)

	f, err := ioutil.ReadFile(file)
	if os.IsNotExist(err) {
		log.Printf("config error: %v", err)
		return nil
	}

	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	if err := json.Unmarshal(f, r); err != nil {
		log.Fatalf("config error: %v", err)
	}

	return r
}

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

	srv := http.NewServer(ic, mux, config(*flagConfig))

	fatal(
		srv.AlertManager,
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
