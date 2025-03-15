.DEFAULT_GOAL:all

.PHONY: all
all: compile build push

.PHONY: compile
compile:
	go build -o bin/payment-rpc-engine .

.PHONY: build
build:
	docker build -t payment-rpc-engine:latest .

.PHONY: push
push:
	docker push payment-rpc-engine:latest