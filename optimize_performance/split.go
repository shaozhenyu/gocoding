package main

import (
	"bytes"
	"strings"
	"testing"
)

func testString() string {
	var s strings.Builder
	for i := 0; i < 1000; i++ {
		s.WriteString("aa")
	}
	return s.String()
}

func testByte() string {
	var b bytes.Buffer
	for i := 0; i < 1000; i++ {
		b.WriteString("aa")
	}
	return b.String()
}

func BenchmarkTestString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testString()
	}
}

func BenchmarkTestByte(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testByte()
	}
}
