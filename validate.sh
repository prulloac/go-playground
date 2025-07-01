#!/bin/bash

# Go Algorithms Playground - Project Validation Script
# This script validates the entire project structure and runs comprehensive tests

echo "🚀 Go Algorithms Playground - Project Validation"
echo "================================================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed or not in PATH"
    exit 1
fi

echo "✅ Go version: $(go version)"

# Check if we're in the right directory
if [ ! -f "go.mod" ]; then
    echo "❌ go.mod not found. Please run this script from the project root."
    exit 1
fi

echo "✅ Project structure validated"

# Initialize modules
echo "📦 Initializing Go modules..."
go mod tidy
if [ $? -ne 0 ]; then
    echo "❌ Failed to initialize Go modules"
    exit 1
fi
echo "✅ Go modules initialized"

# Format code
echo "🎨 Formatting Go code..."
go fmt ./...
echo "✅ Code formatted"

# Run static analysis
echo "🔍 Running static analysis..."
go vet ./...
if [ $? -ne 0 ]; then
    echo "❌ Static analysis failed"
    exit 1
fi
echo "✅ Static analysis passed"

# Run all tests
echo "🧪 Running all tests..."
go test -v ./...
if [ $? -ne 0 ]; then
    echo "❌ Tests failed"
    exit 1
fi
echo "✅ All tests passed"

# Run tests with coverage
echo "📊 Running tests with coverage..."
go test -cover ./...
if [ $? -ne 0 ]; then
    echo "❌ Coverage analysis failed"
    exit 1
fi
echo "✅ Coverage analysis completed"

# Run benchmarks
echo "⚡ Running benchmarks..."
go test -bench=. ./benchmarks
if [ $? -ne 0 ]; then
    echo "❌ Benchmarks failed"
    exit 1
fi
echo "✅ Benchmarks completed"

# Build example
echo "🏗️  Building example application..."
go build -o bin/playground examples/main.go
if [ $? -ne 0 ]; then
    echo "❌ Failed to build example application"
    exit 1
fi
echo "✅ Example application built successfully"

# Run example
echo "🎯 Running example application..."
./bin/playground
if [ $? -ne 0 ]; then
    echo "❌ Example application failed to run"
    exit 1
fi
echo "✅ Example application ran successfully"

# Cleanup
echo "🧹 Cleaning up..."
rm -f bin/playground
echo "✅ Cleanup completed"

echo ""
echo "🎉 Project validation completed successfully!"
echo "✅ All algorithms and data structures are working correctly"
echo "✅ All tests pass"
echo "✅ Code quality checks pass"
echo "✅ Example application works"
echo ""
echo "Your Go Algorithms Playground is ready for development! 🚀"
