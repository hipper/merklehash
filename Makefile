#!/usr/bin/env bash

BINARY_NAME := merklehash

LD_FLAGS=-ldflags " \
    -X main.goos=$(shell go env GOOS) \
    -X main.goarch=$(shell go env GOARCH) \
    -X main.gitCommit=$(shell git rev-parse HEAD) \
    -X main.buildDate=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') \
    "
.PHONY: build
build: ## Build the project locally
	go build $(LD_FLAGS) -o bin/${BINARY_NAME} ./cmd/${BINARY_NAME}

.PHONY: check-fmt
check-fmt:
	@test -z "$(shell gofmt -l -s . | tee /dev/stderr)"

.PHONY: lint
lint: golangci-lint
	$(GOLANGCI_LINT) run

GOLANGCI_LINT = $(shell pwd)/bin/golangci-lint
golangci-lint:
	@[ -f $(GOLANGCI_LINT) ] || { \
	set -e ;\
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell dirname $(GOLANGCI_LINT)) v1.47.2 ;\
	}

.PHONY: test
test: test-unit

.PHONY: test-unit
test-unit:
	go test -race ./...
