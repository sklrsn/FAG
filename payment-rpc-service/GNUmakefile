.DEFAULT_GOAL:all

.PHONY: all
all: compile build

.PHONY: compile
compile:
	go build -o bin/payment-rpc-engine .

.PHONY: build
build:
	docker build -t payment-rpc-engine:latest .