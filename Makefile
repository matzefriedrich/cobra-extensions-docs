.PHONY: help build install-tools lint-fix clean

BIN_DIR := bin
BUILD_DIR := build
GOLANGCI_LINT_VERSION := 2.12.2

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build all sample apps
	@echo "Building sample apps..."
	@mkdir -p $(BUILD_DIR)
	go build -v -o $(BUILD_DIR)/ ./cmd/...

install-tools: ## Install required tools (golangci-lint)
	@echo "Installing golangci-lint v$(GOLANGCI_LINT_VERSION)..."
	@mkdir -p $(BIN_DIR)
	GOBIN=$(shell pwd)/$(BIN_DIR) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v$(GOLANGCI_LINT_VERSION)

lint-fix: ## Run golangci-lint --fix
	@echo "Running golangci-lint --fix..."
	@if [ ! -f $(BIN_DIR)/golangci-lint ]; then \
		echo "golangci-lint not found in $(BIN_DIR). Please run 'make install-tools' first."; \
		exit 1; \
	fi
	$(BIN_DIR)/golangci-lint run --fix

clean: ## Remove bin and build folders
	@echo "Cleaning..."
	rm -rf $(BIN_DIR) $(BUILD_DIR)
