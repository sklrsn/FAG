.DEFAULT_GOAL:all

.PHONY: all
all: compile build

.PHONY: compile
compile:
	go build -o bin/order-rpc-engine .

.PHONY: build
build:
	docker build -t order-rpc-engine:latest .