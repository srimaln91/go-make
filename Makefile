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
BINARY=$(BINARY_NAME)-$(OS)-$(ARCH)

# This version-strategy uses git tags to set the version string
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-list -1 HEAD)

# This version-strategy uses a manual value to set the version string
#VERSION := 1.2.3

BIN_FILE=$(BUILD_DIR)/$(VERSION)/$(BINARY)

.PHONY: all
all: test build ## Test and build

.PHONY: build
build: ## Build the binary
	echo $(BINARY_NAME)
	rm -f $(BIN_FILE)

	CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(ARCH) $(GOBUILD) -ldflags "-X github.com/srimaln91/go-make.version=$(VERSION) \
	-X github.com/srimaln91/go-make.date=$(DATE) \
	-X github.com/srimaln91/go-make.gitCommit=$(COMMIT) \
	-X github.com/srimaln91/go-make.osArch=$(OS)/$(ARCH)" \
	-o $(BIN_FILE) -v

.PHONY: test
test: ## Run unit tests
	$(GOTEST) -covermode=atomic -coverprofile=coverage.txt -v ./... -ldflags "-X github.com/srimaln91/go-make.version=$(VERSION) \
	-X github.com/srimaln91/go-make.date=$(DATE) \
	-X github.com/srimaln91/go-make.gitCommit=$(COMMIT) \
	-X github.com/srimaln91/go-make.osArch=$(OS)/$(ARCH)"

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