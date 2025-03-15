.DEFAULT_GOAL:all

.PHONY: all
all: compile build

.PHONY: compile
compile:
	go build -o bin/shipping-rpc-engine .

.PHONY: build
build:
	docker build -t shipping-rpc-engine:latest .