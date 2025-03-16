FROM --platform=amd64 golang:1.24-bullseye

RUN apt-get update -y && apt-get install -y dnsutils net-tools curl telnet

WORKDIR /go/src/github.com/sklrsn/FAG/
RUN go clean -cache
COPY *.go .
COPY go.mod .
COPY go.sum .
RUN go build -o /opt/bin/shipping-rpc-engine

CMD [ "/opt/bin/shipping-rpc-engine" ]