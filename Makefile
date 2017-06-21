REL = r0
VERSION =
GIT_REVLIST = $(shell test -d .git && git rev-list --count HEAD)
GIT_DESCRIBE = $(shell test -d .git && git describe --always)
GIT_REF = $(shell test -d .git && git rev-parse --abbrev-ref HEAD)

ifneq "$(GIT_DESCRIBE)" ""
REL = r$(GIT_REVLIST).$(GIT_DESCRIBE)
endif

ifneq "$(GIT_REF)" "master"
VERSION = $(REL)-$(GIT_REF)
else
VERSION = $(REL)
endif

default: bot-daemon

all: test bot-daemon bot-client

bot-daemon:
	CGO_ENABLED=0 go build -i -ldflags "-X main.buildversion=$(VERSION)" -v github.com/multimfi/bot/cmd/bot-daemon

bot-client:
	CGO_ENABLED=0 go build -i -ldflags "-X main.buildversion=$(VERSION)" -v github.com/multimfi/bot/cmd/bot-client

install: bot-client
	install bot-client $(HOME)/.local/bin/bot-client

test:
	CGO_ENABLED=1 go test -race ./pkg/...

clean:
	rm -v bot-client bot-daemon

.PHONY: bot-daemon bot-client
