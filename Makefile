.PHONY: build
build:
	go build -v ./cmd/main.go
	./main

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build