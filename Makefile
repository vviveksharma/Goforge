.PHONY: install build clean test lint help

# Default target
.DEFAULT_GOAL := help

# Install goforge to $GOPATH/bin
install:
	@echo "Building and installing goforge..."
	@go build -ldflags="-s -w" -o $(shell go env GOPATH)/bin/goforge ./cmd/goforge
	@echo "✓ goforge installed to $(shell go env GOPATH)/bin/goforge"
	@echo "Make sure $(shell go env GOPATH)/bin is in your PATH"

# Build binary to local directory
build:
	@echo "Building goforge..."
	@go build -o goforge ./cmd/goforge
	@echo "✓ Binary created: ./goforge"

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run linter
lint:
	@echo "Running linters..."
	@go fmt ./...
	@go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@rm -f goforge
	@rm -rf dist/
	@echo "✓ Clean complete"

# Show help
help:
	@echo "GoForge Makefile"
	@echo ""
	@echo "Usage:"
	@echo "  make install    Build and install goforge to \$$GOPATH/bin"
	@echo "  make build      Build goforge binary locally"
	@echo "  make test       Run tests"
	@echo "  make lint       Run formatters and linters"
	@echo "  make clean      Remove build artifacts"
	@echo "  make help       Show this help message"
