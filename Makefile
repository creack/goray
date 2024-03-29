NAME  = goray
SHELL = /bin/bash

SRCS = $(shell find . -name '*.go' -or -name '*.yanl' -or -name '*.json') Makefile Dockerfile

uname_arch := $(shell uname -m)

GOOS   ?= $(shell uname -s | tr '[:upper:]' '[:lower:]')
GOARCH ?= amd64

GO_VERSION = 1.21

ifneq (${uname_arch},x86_64)
GOARCH=${uname_arch}
endif

${NAME}: .built
	docker run --rm ${NAME} cat /tmp/out > ${NAME}
	@chmod +x ${NAME}

.built: ${SRCS}
	docker build -t ${NAME} --build-arg GOOS=${GOOS} --build-arg GOARCH=${GOARCH} --build-arg GO_VERSION=${GO_VERSION} .
	@touch $@

_$(NAME): $(SRCS)
	GO111MODULE=off GOPATH=${GOPATH}:${PWD}/Godeps/_workspace go build -o $@

.PHONY: fmt _fmt
fmt: .built
	docker run --rm ${NAME} make _fmt

_fmt:
	@[ -z "$(shell gofmt -s -l . 2>&1 | \grep -v ^Godeps || true)" ] || (echo "go fmt errors:"; gofmt -s -l . | \grep -v ^Godeps; exit 1)
	@[ -z "$(shell go vet ./...  2>&1 | \grep -v ^Godeps || true)" ] || (echo "go vet errors:"; go vet ./... | \grep -v ^Godeps;  exit 1)
# golint is no more and won't install with go1.5.
#		@[ -z "$(shell golint ./...  2>&1 | \grep -v ^Godeps || true)" ] || (echo "golint errors:"; golint ./...;  exit 1)

.PHONY: clean
clean:
	rm -f .built _${NAME} ${NAME}
