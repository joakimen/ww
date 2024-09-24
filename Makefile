SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

all: fmt lint vet build test

.PHONY: fmt
fmt:
	# github.com/mvdan/gofumpt
	gofumpt -l -w .

.PHONY: lint
lint:
	golangci-lint run ./...

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: build
build:
	go build -v

.PHONY: test
test:
	go test ./...

