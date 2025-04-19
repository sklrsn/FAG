FROM --platform=amd64 golang:1.24-bullseye

RUN apt-get update -y && apt-get install -y dnsutils net-tools curl telnet iputils-ping

WORKDIR /go/src/github.com/sklrsn/FAG/order-stream
RUN go clean -cache
ADD . .

RUN go build -o /opt/bin/order-stream
CMD [ "/opt/bin/order-stream" ]
