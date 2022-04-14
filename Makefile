.PHONY: build
build:
		go build -v ./cmd/apiserver
		docker start my-redis
		./apiserver

.PHONY: test
test:
		go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

redis:
	docker pull redis:latest
	docker run --name my-redis -p 6379:6379 -d redis