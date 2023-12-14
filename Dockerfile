ARG GO_VERSION=1.5
FROM golang:${GO_VERSION} as base

ARG GOOS=linux \
    GOARCH=amd64
ENV GOOS=${GOOS} \
    GOARCH=${GOARCH}

RUN go version

RUN go build -a std

WORKDIR ${GOPATH}/src/github.com/creack/goray
ADD . .

RUN make _goray GO_VERSION=${GO_VERSION} GOOS=${GOOS} GOARCH=${GOARCH} && cp _goray /tmp/out
