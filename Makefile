SHELL := /usr/bin/env bash -o pipefail
.SHELLFLAGS := -ec

# Project variables
BINARY_NAME := ww
BINARY_PATH := bin/$(BINARY_NAME)
BUILD_FLAGS := -v
COVERAGE_DIR := coverage

# Go variables
GO := go
GOFMT := gofumpt
GOLINT := golangci-lint
GOPATH ?= $(shell go env GOPATH)

.DEFAULT_GOAL := help

.PHONY: all
all: clean fmt lint vet test build ## Run all targets

.PHONY: help
help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Remove build artifacts
	rm -rf bin/ coverage/

.PHONY: fmt
fmt: ## Format code using gofumpt
	$(GOFMT) -l -w .

.PHONY: lint
lint: ## Run golangci-lint
	$(GOLINT) run ./...

.PHONY: lint-fix
lint-fix: ## Run golangci-lint with auto-fix
	$(GOLINT) run --fix ./...

.PHONY: vet
vet: ## Run go vet
	$(GO) vet ./...

.PHONY: build
build: ## Build binary
	mkdir -p bin
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_PATH)

.PHONY: test
test: ## Run tests
	$(GO) test -v ./...

.PHONY: coverage
coverage: ## Run tests with coverage
	mkdir -p $(COVERAGE_DIR)
	$(GO) test -v -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	$(GO) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html

.PHONY: install
install: build ## Install binary to $GOPATH/bin
	mkdir -p $(GOPATH)/bin
	cp $(BINARY_PATH) $(GOPATH)/bin/

.PHONY: run
run: build ## Run the application
	./$(BINARY_PATH)
