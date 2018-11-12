
FROM golang:latest

RUN go get github.com/dgrijalva/jwt-go "github.com/utahta/go-cronowriter"

COPY . /go/src/api_gateway/gateway

WORKDIR /go/src/api_gateway/gateway

RUN go install api_gateway/gateway

ENTRYPOINT /go/bin/gateway
