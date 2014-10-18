FROM google/golang:stable
ADD  .	/gopath/src/github.com/creack/goray
RUN  cd /gopath/src/github.com/creack/goray && go get -d
WORKDIR /gopath/src/github.com/creack/goray
