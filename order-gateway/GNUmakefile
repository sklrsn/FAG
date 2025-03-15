.DEFAULT_GOAL:all

.PHONY: all
all: compile build push

.PHONY: compile
compile:
	go build -o bin/order-gateway .

.PHONY: build
build:
	docker build -t order-gateway:latest .

.PHONY: push
push:
	docker push order-gateway:latest