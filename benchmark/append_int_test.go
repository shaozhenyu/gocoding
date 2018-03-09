package main

import (
	"testing"
)

func main() {
}

func array() [1000000]int {
	var x [1000000]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice(count int) []int {
	x := make([]int, count)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice1(count int) []int {
	x := []int{}
	for i := 0; i < count; i++ {
		x = append(x, i)
	}
	return x
}



func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice(10000000)
	}
}

func BenchmarkSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1(10000000)
	}
}
