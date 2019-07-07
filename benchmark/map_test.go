package main

import (
	"testing"
	"fmt"
)

func str(times int) {
	m := make(map[string]bool, times)
	for i := 0; i < times; i++ {
		m[fmt.Sprintf("%d+%d", times, times)] =  true
	}
}

func arr(times int) {
	m := make(map[[2]int]bool, times)
	for i := 0; i < times; i++ {
		m[[2]int{times, times}] =  true
	}
}

func BenchmarkStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str(100000)
	}
}

func BenchmarkArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr(100000)
	}
}
