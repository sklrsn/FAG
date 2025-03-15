.DEFAULT_GOAL:all

.PHONY: all
all: compile build

.PHONY: compile
compile:
	go build -o bin/order-gateway .

.PHONY: build
build:
	docker build -t order-gateway:latest .