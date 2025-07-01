package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"random order", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
		{"duplicates", []int{3, 3, 1, 1, 2, 2}, []int{1, 1, 2, 2, 3, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := BubbleSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("BubbleSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"random order", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := QuickSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("QuickSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reverse sorted", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"random order", []int{3, 1, 4, 1, 5, 9, 2, 6}, []int{1, 1, 2, 3, 4, 5, 6, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MergeSort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
