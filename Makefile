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
all: clean fmt lint vet test build install ## Run all targets

.PHONY: help
help: ## Show this help
	@echo "=== Running target: help ==="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## Remove build artifacts
	@echo "=== Running target: clean ==="
	rm -rf bin/ coverage/

.PHONY: fmt
fmt: ## Format code using gofumpt
	@echo "=== Running target: fmt ==="
	$(GOFMT) -l -w .

.PHONY: lint
lint: ## Run golangci-lint
	@echo "=== Running target: lint ==="
	$(GOLINT) run ./...

.PHONY: lint-fix
lint-fix: ## Run golangci-lint with auto-fix
	@echo "=== Running target: lint-fix ==="
	$(GOLINT) run --fix ./...

.PHONY: vet
vet: ## Run go vet
	@echo "=== Running target: vet ==="
	$(GO) vet ./...

.PHONY: build
build: ## Build binary
	@echo "=== Running target: build ==="
	mkdir -p bin
	$(GO) build $(BUILD_FLAGS) -o $(BINARY_PATH)

.PHONY: test
test: ## Run tests
	@echo "=== Running target: test ==="
	$(GO) test -v ./...

.PHONY: coverage
coverage: ## Run tests with coverage
	@echo "=== Running target: coverage ==="
	mkdir -p $(COVERAGE_DIR)
	$(GO) test -v -coverprofile=$(COVERAGE_DIR)/coverage.out ./...
	$(GO) tool cover -html=$(COVERAGE_DIR)/coverage.out -o $(COVERAGE_DIR)/coverage.html

.PHONY: install
install: build ## Install binary to $GOPATH/bin
	@echo "=== Running target: install ==="
	mkdir -p $(GOPATH)/bin
	cp $(BINARY_PATH) $(GOPATH)/bin/

.PHONY: run
run: build ## Run the application
	@echo "=== Running target: run ==="
	./$(BINARY_PATH)
