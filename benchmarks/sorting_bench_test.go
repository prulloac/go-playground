package benchmarks

import (
	"fmt"
	"testing"

	"github.com/prulloac/go-playground/algorithms/sorting"
	"github.com/prulloac/go-playground/utils"
)

func BenchmarkBubbleSort(b *testing.B) {
	sizes := []int{100, 1000, 5000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			data := utils.GenerateRandomSlice(size, 1, 1000)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				testData := utils.CopySlice(data)
				sorting.BubbleSort(testData)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	sizes := []int{100, 1000, 5000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			data := utils.GenerateRandomSlice(size, 1, 1000)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				testData := utils.CopySlice(data)
				sorting.QuickSort(testData)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	sizes := []int{100, 1000, 5000, 10000}

	for _, size := range sizes {
		b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
			data := utils.GenerateRandomSlice(size, 1, 1000)
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				testData := utils.CopySlice(data)
				sorting.MergeSort(testData)
			}
		})
	}
}
