.PHONY: build test install clean run

APP_NAME=tk
CMD_PATH=./cmd/toolkit

# Default target
all: build

# Build the local ecosystem at the project root
build:
	@echo "Building binary $(APP_NAME)..."
	go build -o $(APP_NAME) $(CMD_PATH)

# Run unit tests and basic linting
test:
	@echo "Running unit test suite..."
	go test ./... -v -count=1

# Install app into global GOPATH/bin
install:
	@echo "Installing $(APP_NAME) to GOPATH/bin..."
	go install $(CMD_PATH)

# Clean build artifacts
clean:
	@echo "Cleaning artifacts..."
	rm -f $(APP_NAME) $(APP_NAME).exe

# Execute CLI in fast-compile mode (dev workflow)
run:
	@go run $(CMD_PATH)/main.go
