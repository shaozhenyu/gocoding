package main

import (
	"strings"
	"unsafe"
	"testing"
)

var s = strings.Repeat("a", 1024)

func main() {
}

func byteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func stringToByte(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func test() {
	b := []byte(s)
	_ = string(b)
}

func test1() {
	b := stringToByte(s)
	_ = byteToString(b)
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}

func BenchmarkTest1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test1()
	}
}
