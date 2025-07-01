# Go Algorithms Playground Makefile

.PHONY: all build test bench clean fmt vet help run

# Default target
all: fmt vet test

# Build the example application
build:
	@echo "Building example application..."
	go build -o bin/playground examples/main.go

# Run all tests
test:
	@echo "Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	go test -cover ./...
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run benchmarks
bench:
	@echo "Running benchmarks..."
	go test -bench=. -benchmem ./benchmarks

# Format Go code
fmt:
	@echo "Formatting Go code..."
	go fmt ./...

# Run static analysis
vet:
	@echo "Running go vet..."
	go vet ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f bin/playground
	rm -f coverage.out coverage.html

# Run the example application
run: build
	@echo "Running example application..."
	./bin/playground

# Initialize Go modules
mod-init:
	@echo "Initializing Go modules..."
	go mod tidy

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download

# Run a specific test package
test-sorting:
	@echo "Running sorting tests..."
	go test -v ./algorithms/sorting

test-searching:
	@echo "Running searching tests..."
	go test -v ./algorithms/searching

test-linkedlist:
	@echo "Running linked list tests..."
	go test -v ./data_structures/linked_list

# Performance profiling
profile-cpu:
	@echo "Running CPU profiling..."
	go test -bench=. -cpuprofile=cpu.prof ./benchmarks
	go tool pprof cpu.prof

profile-mem:
	@echo "Running memory profiling..."
	go test -bench=. -memprofile=mem.prof ./benchmarks
	go tool pprof mem.prof

# Help target
help:
	@echo "Available targets:"
	@echo "  all          - Run fmt, vet, and test"
	@echo "  build        - Build the example application"
	@echo "  test         - Run all tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  bench        - Run benchmarks"
	@echo "  fmt          - Format Go code"
	@echo "  vet          - Run static analysis"
	@echo "  clean        - Clean build artifacts"
	@echo "  run          - Build and run example application"
	@echo "  mod-init     - Initialize Go modules"
	@echo "  deps         - Install dependencies"
	@echo "  profile-cpu  - Run CPU profiling"
	@echo "  profile-mem  - Run memory profiling"
	@echo "  help         - Show this help message"
