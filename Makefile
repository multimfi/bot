CGO = 0
VERSION = r0
GIT_REVLIST = $(shell test -d .git && git rev-list --count HEAD)
GIT_DESCRIBE = $(shell test -d .git && git describe --always)

ifneq "$(GIT_DESCRIBE)" ""
VERSION = r$(GIT_REVLIST).$(GIT_DESCRIBE)
endif

default:
	CGO_ENABLED=$(CGO) go build -i -ldflags "-X main.buildversion=$(VERSION)" -v -o $(PWD)/bot .
