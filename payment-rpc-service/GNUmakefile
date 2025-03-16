.DEFAULT_GOAL:all

.PHONY: all
all: build push

.PHONY: build
build:
	docker build -t payment-rpc-engine:latest .

.PHONY: push
push:
	docker tag payment-rpc-engine:latest sklrsn/payment-rpc-engine:latest
	docker push sklrsn/payment-rpc-engine:latest