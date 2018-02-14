REL = r0

GIT_DESCRIBE := $(shell test -d .git && git describe --always)
GIT_REVLIST := $(shell test -d .git && git rev-list --count HEAD)
GIT_DIRTY := $(shell test -d .git && git diff-index --quiet HEAD || date '+-dirty-%s')
GIT_REF := $(shell test -d .git && git rev-parse --abbrev-ref HEAD)

ifneq "$(GIT_DESCRIBE)" ""
REL = r$(GIT_REVLIST).$(GIT_DESCRIBE)
endif

ifndef VERSION
ifneq "$(GIT_REF)" "master"
VERSION = $(REL)-$(GIT_REF)$(GIT_DIRTY)
else
VERSION = $(REL)$(GIT_DIRTY)
endif
endif

ifndef BUILDFLAGS
BUILDFLAGS = -i -v
endif

ifndef PKGS
PKGS := $(shell go list -f '{{if .TestGoFiles}}{{.ImportPath}}{{end}}' ./... | grep -v /vendor/)
endif

default: bot-daemon

all: static test bot-daemon bot-client

bot-daemon bot-client:
	go build -tags "$(GOTAGS)" -ldflags "-X main.buildversion=$(VERSION)" $(BUILDFLAGS) github.com/multimfi/bot/cmd/$@

install: bot-client
	install bot-client $(HOME)/.local/bin/bot-client

static:
	$(MAKE) -C static

test-install:
	CGO_ENABLED=1 go test -race -i -v $(PKGS)

test:
	CGO_ENABLED=1 go test -race $(PKGS)

clean:
	rm -v bot-client bot-daemon

.PHONY: bot-daemon bot-client static
