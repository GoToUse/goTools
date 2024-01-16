package goTools

import (
	"testing"
)

var arrays = [][]int{
	[]int{1, 2, 3},
	[]int{4, 5, 6},
	[]int{7, 8, 9},
	[]int{10, 11, 12},
}

var B []int

func BenchmarkMerge(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = Merge(arrays...)
	}
}

func BenchmarkConcatCopyPreAllocate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		B = ConcatCopyPreAllocate(arrays)
	}
}
