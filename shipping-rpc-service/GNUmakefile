.DEFAULT_GOAL:all

.PHONY: all
all: compile build push

.PHONY: compile
compile:
	go build -o bin/shipping-rpc-engine .

.PHONY: build
build:
	docker build -t shipping-rpc-engine:latest .

.PHONY: push
push:
	docker push shipping-rpc-engine:latest