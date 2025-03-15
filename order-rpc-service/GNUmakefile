.DEFAULT_GOAL:all

.PHONY: all
all: compile build push

.PHONY: compile
compile:
	go build -o bin/order-rpc-engine .

.PHONY: build
build:
	docker build -t order-rpc-engine:latest .

.PHONY: push
push:
	docker tag order-rpc-engine:latest sklrsn/order-rpc-engine:latest
	docker push sklrsn/order-rpc-engine:latest