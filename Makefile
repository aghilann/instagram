BINARY=bin/main

build:
	@echo "Building the Go application..."
	go build -o $(BINARY) cmd/main.go

# Run the Go application
run: build
	@echo "Running the Go application..."
	./$(BINARY)

# Run all Go tests
test:
	@echo "Running Go tests..."
	go test ./...

# Clean up the build binary
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY)

# Phony targets to avoid file conflicts
.PHONY: build run test clean
