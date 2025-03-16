.DEFAULT_GOAL:all

.PHONY: all
all: build push

.PHONY: build
build:
	docker build -t order-gateway:latest .

.PHONY: push
push:
	docker tag order-gateway:latest sklrsn/order-gateway:latest
	docker push sklrsn/order-gateway:latest