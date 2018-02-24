FROM golang:1.9.3-alpine

WORKDIR /go/src/github.com/matsu-chara/conbot

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git

RUN go get github.com/golang/dep/cmd/dep

COPY Gopkg.toml .
COPY Gopkg.lock .
RUN dep ensure -vendor-only

COPY . .
RUN go build && go install

ENTRYPOINT ["conbot"]
