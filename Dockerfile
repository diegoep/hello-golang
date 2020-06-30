FROM golang

ADD . /go/src/main

RUN go get github.com/gin-gonic/gin
RUN go install main

ENTRYPOINT /go/bin/main

EXPOSE 8080