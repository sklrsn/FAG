.DEFAULT_GOAL:all

.PHONY: all
all: build push

.PHONY: build
build:
	docker build -t order-stream:latest .

.PHONY: push
push:
	docker tag order-stream:latest sklrsn/order-stream:latest
	docker push sklrsn/order-stream:latest

.PHONY: gqlgen-generate
gqlgen-generate:
	gqlgen generate

.PHONY: gqlgen-init
gqlgen-init:
	gqlgen init