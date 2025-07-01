package utils

import (
	"math/rand"
)

// GenerateRandomSlice creates a slice of random integers
func GenerateRandomSlice(size int, min, max int) []int {
	slice := make([]int, size)

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(max-min+1) + min
	}

	return slice
}

// IsSorted checks if a slice is sorted in ascending order
func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// CopySlice creates a deep copy of an integer slice
func CopySlice(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	return result
}

// Min returns the minimum value in a slice
func Min(arr []int) int {
	if len(arr) == 0 {
		panic("cannot find minimum of empty slice")
	}

	min := arr[0]
	for _, val := range arr[1:] {
		if val < min {
			min = val
		}
	}
	return min
}

// Max returns the maximum value in a slice
func Max(arr []int) int {
	if len(arr) == 0 {
		panic("cannot find maximum of empty slice")
	}

	max := arr[0]
	for _, val := range arr[1:] {
		if val > max {
			max = val
		}
	}
	return max
}

// Sum returns the sum of all elements in a slice
func Sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// Average returns the average of all elements in a slice
func Average(arr []int) float64 {
	if len(arr) == 0 {
		return 0
	}
	return float64(Sum(arr)) / float64(len(arr))
}
