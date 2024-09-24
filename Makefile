SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

all: build fmt vet lint test

.PHONY: build
build:
	go build -v

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix ./...
