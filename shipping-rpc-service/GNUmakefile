.DEFAULT_GOAL:all

.PHONY: all
all: build push

.PHONY: build
build:
	docker build -t shipping-rpc-engine:latest .

.PHONY: push
push:
	docker tag shipping-rpc-engine:latest sklrsn/shipping-rpc-engine:latest
	docker push sklrsn/shipping-rpc-engine:latest