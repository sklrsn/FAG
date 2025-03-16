FROM golang:1.24-bullseye

WORKDIR /go/src/github.com/sklrsn/FAG/
COPY *.go .
COPY go.mod .
COPY go.sum .
RUN go build -o /opt/bin/order-gateway

CMD [ "/opt/bin/order-gateway" ]