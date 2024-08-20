GOPATH:=$(shell go env GOPATH)

.PHONY: run
run:
	go run ./cmd/app/main.go

.PHONY: build
build:
	go build ./cmd/app/main.go

.PHONY: tidy
tidy:
	go mod tidy