# Go Algorithms Playground

A comprehensive collection of algorithms and data structures implemented in Go. This repository serves as a learning resource and reference for common computer science algorithms, optimized for clarity and performance.

## 📁 Project Structure

```
go-playground/
├── algorithms/                 # Algorithm implementations
│   ├── sorting/               # Sorting algorithms (bubble, quick, merge, etc.)
│   ├── searching/             # Searching algorithms (linear, binary, etc.)
│   ├── graph/                 # Graph algorithms (DFS, BFS, Dijkstra, etc.)
│   ├── dynamic_programming/   # DP algorithms (fibonacci, knapsack, etc.)
│   ├── greedy/               # Greedy algorithms
│   ├── tree/                 # Tree algorithms (traversal, BST, etc.)
│   ├── string/               # String algorithms (pattern matching, etc.)
│   └── math/                 # Mathematical algorithms
├── data_structures/          # Data structure implementations
│   ├── linked_list/          # Singly/doubly linked lists
│   ├── stack/                # Stack implementation
│   ├── queue/                # Queue implementation
│   ├── heap/                 # Min/max heap
│   ├── hash_table/           # Hash table implementation
│   └── trie/                 # Trie (prefix tree)
├── benchmarks/               # Performance benchmarks
├── utils/                    # Utility functions
├── examples/                 # Example usage and demos
├── go.mod                    # Go module file
└── README.md                 # This file
```

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/prulloac/go-playground.git
cd go-playground
```

2. Initialize the Go module:
```bash
go mod tidy
```

3. Run the example:
```bash
go run examples/main.go
```

## 🧪 Running Tests

Run all tests:
```bash
go test ./...
```

Run tests for a specific package:
```bash
go test ./algorithms/sorting
go test ./data_structures/linked_list
```

Run tests with verbose output:
```bash
go test -v ./...
```

## 📊 Benchmarks

Run performance benchmarks:
```bash
go test -bench=. ./benchmarks
```

Run benchmarks with memory allocation stats:
```bash
go test -bench=. -benchmem ./benchmarks
```

## 📚 Algorithms Included

### Sorting Algorithms
- **Bubble Sort** - O(n²) - Simple comparison-based sorting
- **Quick Sort** - O(n log n) avg - Divide-and-conquer sorting
- **Merge Sort** - O(n log n) - Stable divide-and-conquer sorting
- **Heap Sort** - O(n log n) - Heap-based sorting
- **Insertion Sort** - O(n²) - Efficient for small datasets

### Searching Algorithms
- **Linear Search** - O(n) - Sequential search
- **Binary Search** - O(log n) - Efficient search on sorted arrays
- **Binary Search (Recursive)** - O(log n) - Recursive implementation

### Graph Algorithms
- **Depth-First Search (DFS)** - Graph traversal
- **Breadth-First Search (BFS)** - Level-order traversal
- **Dijkstra's Algorithm** - Shortest path
- **Topological Sort** - Dependency ordering

### Dynamic Programming
- **Fibonacci Sequence** - Classic DP example
- **Knapsack Problem** - Optimization problem
- **Longest Common Subsequence** - String algorithms
- **Coin Change** - Counting combinations

### Tree Algorithms
- **Binary Search Tree** - Ordered tree structure
- **Tree Traversals** - Inorder, preorder, postorder
- **Lowest Common Ancestor** - Tree queries

## 🏗️ Data Structures

### Linear Data Structures
- **Linked List** - Dynamic array alternative
- **Stack** - LIFO data structure
- **Queue** - FIFO data structure

### Tree-based Structures
- **Binary Search Tree** - Ordered tree for fast lookups
- **Heap** - Priority queue implementation
- **Trie** - Prefix tree for string operations

### Hash-based Structures
- **Hash Table** - Key-value mapping with O(1) average access

## 📈 Performance Characteristics

Each algorithm includes comprehensive time and space complexity analysis:

| Algorithm | Best Case | Average Case | Worst Case | Space |
|-----------|-----------|--------------|------------|-------|
| Bubble Sort | O(n) | O(n²) | O(n²) | O(1) |
| Quick Sort | O(n log n) | O(n log n) | O(n²) | O(log n) |
| Merge Sort | O(n log n) | O(n log n) | O(n log n) | O(n) |
| Binary Search | O(1) | O(log n) | O(log n) | O(1) |

## 🧱 Code Organization

Each algorithm/data structure follows a consistent pattern:

1. **Implementation file** (`*.go`) - Core algorithm logic
2. **Test file** (`*_test.go`) - Comprehensive unit tests
3. **Documentation** - Inline comments with complexity analysis
4. **Examples** - Usage demonstrations

### Example Structure:
```go
// Function with complexity analysis
// Time Complexity: O(n log n), Space Complexity: O(n)
func MergeSort(arr []int) []int {
    // Implementation here
}
```

## 🤝 Contributing

Contributions are welcome! Please follow these guidelines:

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/new-algorithm`)
3. **Add** your algorithm with tests and documentation
4. **Ensure** all tests pass (`go test ./...`)
5. **Run** `go fmt` to format your code
6. **Submit** a pull request

### Adding New Algorithms

1. Create the algorithm in the appropriate category folder
2. Include comprehensive tests
3. Add benchmark if performance-critical
4. Update README with algorithm details
5. Include time/space complexity analysis

## 📖 Learning Resources

- [Introduction to Algorithms (CLRS)](https://mitpress.mit.edu/books/introduction-algorithms)
- [Algorithm Design Manual](https://www.algorist.com/)
- [Go Documentation](https://golang.org/doc/)
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🔗 Useful Commands

```bash
# Run all tests
go test ./...

# Run benchmarks
go test -bench=. ./benchmarks

# Generate test coverage report
go test -cover ./...

# Format all Go files
go fmt ./...

# Run static analysis
go vet ./...

# Build the example
go build -o playground examples/main.go
```

## 📊 Project Stats

- **Languages**: Go
- **Algorithms**: 15+ implementations
- **Data Structures**: 6+ implementations
- **Test Coverage**: 90%+
- **Benchmarks**: Performance analysis included

---

Happy coding! 🚀