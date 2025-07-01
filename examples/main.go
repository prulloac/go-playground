package main

import (
	"fmt"

	"github.com/prulloac/go-playground/algorithms/sorting"
	"github.com/prulloac/go-playground/utils"
)

func main() {
	fmt.Println("=== Go Algorithms Playground ===")
	fmt.Println()

	// Demonstrate sorting algorithms
	fmt.Println("1. Sorting Algorithms Demo:")
	unsorted := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Original array: %v\n", unsorted)

	bubbleSorted := sorting.BubbleSort(unsorted)
	fmt.Printf("Bubble Sort:    %v\n", bubbleSorted)

	quickSorted := sorting.QuickSort(unsorted)
	fmt.Printf("Quick Sort:     %v\n", quickSorted)

	mergeSorted := sorting.MergeSort(unsorted)
	fmt.Printf("Merge Sort:     %v\n", mergeSorted)

	// Demonstrate utility functions
	fmt.Println("\n4. Utility Functions Demo:")
	randomData := utils.GenerateRandomSlice(10, 1, 100)
	fmt.Printf("Random data: %v\n", randomData)
	fmt.Printf("Is sorted: %t\n", utils.IsSorted(randomData))
	fmt.Printf("Min: %d, Max: %d\n", utils.Min(randomData), utils.Max(randomData))
	fmt.Printf("Sum: %d, Average: %.2f\n", utils.Sum(randomData), utils.Average(randomData))
}
