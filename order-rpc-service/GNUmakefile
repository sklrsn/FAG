.DEFAULT_GOAL:all

.PHONY: all
all: build push

.PHONY: build
build:
	docker build -t order-rpc-engine:latest .

.PHONY: push
push:
	docker tag order-rpc-engine:latest sklrsn/order-rpc-engine:latest
	docker push sklrsn/order-rpc-engine:latest