FROM golang:1.24-bullseye

WORKDIR /go/src/github.com/sklrsn/FAG/
RUN go clean -cache
COPY *.go .
COPY go.mod .
COPY go.sum .
RUN go build -o /opt/bin/order-rpc-engine

CMD [ "/opt/bin/order-rpc-engine" ]