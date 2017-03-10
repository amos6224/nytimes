FROM golang
MAINTAINER 'jeff destine' amos6224@gmail.como
RUN mkdir -p $GOPATH/src/github.com/amos6224/hello
ADD ./main.go /go/src/github.com/amos6224/hello
RUN go install github.com/amos6224/hello
ENTRYPOINT /go/bin/hello
EXPOSE 8080
