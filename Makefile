.PHONY: build test lint fmt clean run deps

# Build the application
build:
	go build ./cmd/server

# Run tests
test:
	go test -v ./...

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run performance tests with k6
perf:
	@if command -v k6 >/dev/null 2>&1; then \
		k6 run tests/performance/scripts/basic_load_test.js; \
	else \
		echo "k6 not found. Install with: brew install k6 (macOS) or https://k6.io/"; \
	fi

# Format Go code
fmt:
	go fmt ./...

# Lint Go code
lint:
	~/go/bin/golangci-lint run ./...

# Install dependencies
deps:
	go mod tidy
	go mod download

# Run the application
run:
	go run cmd/server/main.go

# Clean build artifacts
clean:
	rm -f coverage.out coverage.html
	go clean -cache

# Install development tools
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# All checks
check: fmt lint test

# Development setup
setup: deps install-tools

# Default target
all: build test