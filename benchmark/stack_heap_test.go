package main

import (
	"testing"
)

func heap() []int {
	x := make([]int, 8193)
	for i := 0; i < 8191; i++ {
		x[i] = i
	}
	return x
}

func stack() []int {
	x := make([]int, 8191)
	for i := 0; i < 8191; i++ {
		x[i] = i
	}
	return x
}


func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		heap()
	}
}

func BenchmarkStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stack()
	}
}
