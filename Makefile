NAME	=	goray
SHELL	=	/bin/bash

SRCS	=	$(shell find . -name '*.go')

$(NAME)	:	$(SRCS)
		go build && go install

fmt		:	$(SRCS)
		@[ -z "$(shell gofmt -s -l . 2>&1 | \grep -v ^Godeps || true)" ] || (echo "go fmt errors:"; gofmt -s -l .; exit 1)
		@[ -z "$(shell golint ./...  2>&1 | \grep -v ^Godeps || true)" ] || (echo "golint errors:"; golint ./...;  exit 1)
		@[ -z "$(shell go vet ./...  2>&1 | \grep -v ^Godeps || true)" ] || (echo "go vet errors:"; go vet ./...;  exit 1)
