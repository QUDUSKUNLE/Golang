# /Dockerfile

FROM golang:alpine

ADD ./src  /go/src/app

WORKDIR /go/src/app

ENV PORT=3000

CMD [ "go", "run", "server.go" ]
