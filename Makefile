# Generic parameters
DATE=$(shell date -u +%Y-%m-%d-%H:%M:%S-%Z)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

BUILD_DIR=build
WORKING_DIRECTORY=$(dir $(realpath $(firstword $(MAKEFILE_LIST))))
BINARY_NAME=$(shell basename $(WORKING_DIRECTORY))

OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
BINARY_LINUX=$(BINARY_NAME)-$(OS)-$(ARCH)

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)
#
# This version-strategy uses a manual value to set the version string
#VERSION := 1.2.3

BIN_FILE=$(BUILD_DIR)/$(VERSION)/$(BINARY_LINUX)

.PHONY: all
all: test build ## Test and build

.PHONY: build
build: ## Build the binary
	echo $(BINARY_NAME)
	rm -f $(BIN_FILE)
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 $(GOBUILD) -ldflags "-X github.com/srimaln91/go-build/util/build.version=$(VERSION) -X github.com/srimaln91/go-build/util/build.date=$(DATE)" -o $(BIN_FILE) -v

.PHONY: test
test: ## Run unit tests
	$(GOTEST) -v ./...

.PHONY: clean
clean: ## Clean the build directory
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

.PHONY: run
run: ## Build and run the binary
	$(GOBUILD) -o $(BIN_FILE) -v
	./$(BIN_FILE)

.PHONY: help
help: ## display help page
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'