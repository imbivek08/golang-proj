APP_NAME := rest-api
GO := go

.PHONY: all build run test clean tidy

all: build

## Build the application
build:
	@echo "🚀 Building $(APP_NAME)..."
	$(GO) build  -o bin/$(APP_NAME) ./cmd

## Run the application
run:
	@echo "🏃 Running $(APP_NAME)..."
	$(GO) run ./cmd/main.go

## Run tests
test:
	@echo "🧪 Running tests..."
	$(GO) test ./... -v

## Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	rm -rf bin

## Format code
fmt:
	@echo "✨ Formatting..."
	$(GO) fmt ./...

## Download dependencies
tidy:
	@echo "📦 Tidying modules..."
	$(GO) mod tidy
