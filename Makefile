# Makefile for the Go static blog generator

# Project name
PROJECT_NAME := blog_generator

# Go binary output directory
BIN_DIR := bin
BINARY_NAME := $(BIN_DIR)/$(PROJECT_NAME)

# Source code directory
SRC_DIR := ./

# Go files to compile
GO_FILES := $(shell find $(SRC_DIR)/cmd -name '*.go' -o -name '*.go' -path '$(SRC_DIR)/internal/*')

# Build directory for the generated website
BUILD_DIR := build

# Default target
.DEFAULT_GOAL := all

# Target to build the Go binary
build:
	@mkdir -p $(BIN_DIR)
	@echo "Building $(BINARY_NAME)..."
	go build -o $(BINARY_NAME) $(GO_FILES)
	@echo "Build complete!"

# Target to generate the static website
generate: build
	@echo "Generating static website to $(BUILD_DIR)..."
	$(BINARY_NAME) # Run the compiled binary
	@echo "Website generation complete!"

# Target to clean up generated files
clean:
	@echo "Cleaning up..."
	rm -rf $(BIN_DIR) $(BUILD_DIR)
	@echo "Clean complete!"

# Target to run the application
run: build generate
	@echo "Running the blog generator..."
	$(BINARY_NAME)
    # No need to re-run generate, it's a dependency.

# Target for building and running
all: build generate

# Install the binary to $GOPATH/bin
install: build
	@echo "Installing $(BINARY_NAME) to $$GOPATH/bin..."
	go install $(GO_FILES) # Corrected: install the package, not the files
	@echo "Installation complete!"

# Target to format the code
fmt:
	@echo "Formatting Go code..."
	go fmt $(GO_FILES)
	@echo "Formatting complete!"

# Target to run tests (if you add them)
test:
	@echo "Running tests..."
	go test ./...
	@echo "Tests complete!"

# Watch for changes and rebuild.  Uses entr if available.
watch:
	@echo "Watching for changes and rebuilding..."
	@which entr >/dev/null 2>&1 || { echo "entr is not installed.  Please install it: brew install entr or https://github.com/gohugoio/entr"; exit 1; }
	find . -name '*.go' -not -path "./build/*" -not -path "./bin/*" | entr -r make build generate

.PHONY: all build clean generate run install fmt test watch

