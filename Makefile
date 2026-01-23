# Copyright 2025 Yaacov Zamir <kobi.zamir@gmail.com>
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

VERSION_GIT := $(shell git describe --tags 2>/dev/null || echo "0.0.0-dev")
VERSION ?= ${VERSION_GIT}

GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GOPATH := $(shell go env GOPATH)

# Source files
PKG_FILES := $(shell find pkg -name '*.go')
CMD_FILES := $(shell find cmd/kubectl-mtv-mcp -name '*.go')
SERVER_FILES := main.go $(CMD_FILES)

# Build flags
LDFLAGS := -s -w -X github.com/yaacov/kubectl-mtv-mcp/cmd/kubectl-mtv-mcp.Version=${VERSION}
BUILD_FLAGS := -trimpath

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo "kubectl-mtv-mcp Makefile"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@awk 'BEGIN {FS = ":.*##"} /^[a-zA-Z_-]+:.*##/ {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: clean ## Build the binary for current platform
	@mkdir -p bin
	@CGO_ENABLED=0 go build $(BUILD_FLAGS) -ldflags='$(LDFLAGS)' -o bin/kubectl-mtv-mcp .
	@echo "Built bin/kubectl-mtv-mcp ($(GOOS)/$(GOARCH))"

.PHONY: install
install: build ## Install binary to GOPATH/bin
	@cp bin/kubectl-mtv-mcp $(GOPATH)/bin/
	@echo "Installed to $(GOPATH)/bin/kubectl-mtv-mcp"

.PHONY: run
run: build ## Build and run the MCP server
	@./bin/kubectl-mtv-mcp

.PHONY: test
test: ## Run tests with coverage
	@go test -v -cover ./pkg/... ./cmd/...

.PHONY: lint
lint: ## Run linters (go vet + golangci-lint)
	@go vet ./pkg/... ./cmd/...
	@$(GOPATH)/bin/golangci-lint run ./pkg/... ./cmd/...

.PHONY: fmt
fmt: ## Format Go source files
	@go fmt ./pkg/... ./cmd/...

.PHONY: tidy
tidy: ## Tidy Go modules
	@go mod tidy

.PHONY: verify
verify: fmt lint test ## Run fmt, lint, and test

.PHONY: clean
clean: ## Remove build artifacts
	@rm -rf bin/ dist/ coverage.out

.PHONY: install-tools
install-tools: ## Install development tools (golangci-lint)
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

.PHONY: dist
dist: ## Build release archive for linux-amd64
	@mkdir -p dist
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $(BUILD_FLAGS) -ldflags='$(LDFLAGS)' -o dist/kubectl-mtv-mcp .
	@tar -C dist -czvf dist/kubectl-mtv-mcp-$(VERSION)-linux-amd64.tar.gz kubectl-mtv-mcp
	@cd dist && sha256sum kubectl-mtv-mcp-$(VERSION)-linux-amd64.tar.gz > kubectl-mtv-mcp-$(VERSION)-linux-amd64.tar.gz.sha256sum
	@rm dist/kubectl-mtv-mcp
	@echo "Built dist/kubectl-mtv-mcp-$(VERSION)-linux-amd64.tar.gz"
