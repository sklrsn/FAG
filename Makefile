.DEFAULT_GOAL: all

.PHONY: all
all: up

.PHONY: push
push:
	(cd order-gateway && make)
	(cd order-rpc-service && make)
	(cd payment-rpc-service && make)
	(cd shipping-rpc-service && make)

.PHONY: up
up:
	docker compose -f docker-compose.yaml up --build

.PHONY: commit
commit: 
	git add .
	git commit -am "fixes from date:$(shell date)"
	git push

generate-cert: clean
	mkdir -p ./certs
	openssl req -x509 -newkey rsa:4096 -keyout certs/nginx.key -out certs/nginx.crt -days 30 \
	-nodes -subj "/CN=www.jaegar.sklrsn.in" -addext "subjectAltName=DNS:jaegar.sklrsn.in,DNS:www.jaegar.sklrsn.in"

clean:
	@rm -rf ./certs