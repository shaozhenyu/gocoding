package main

import (
	"fmt"
	"io"
)

type alphaReader struct {
	src string
	cur int
}

func newAlphaReader(src string) *alphaReader {
	return &alphaReader{src: src}
}

func alpha(r byte) byte {
	if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
		return r
	}
	return 0
}

func (a *alphaReader) Read(p []byte) (int, error) {
	if a.cur >= len(a.src) {
		return 0, io.EOF
	}

	x := len(a.src) - a.cur
	n, bound := 0, 0
	if x >= len(p) {
		bound = len(p)
	} else if x <= len(p) {
		bound = x
	}

	buffer := make([]byte, bound)
	for n < bound {
		if char := alpha(a.src[a.cur]); char != 0 {
			buffer[n] = char
		}
		a.cur++
		n++
	}
	copy(p, buffer)
	return n, nil
}

func main() {
	reader := newAlphaReader("Hello! It's 9am, where is the sun?")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Print(string(p[:n]))
	}
	fmt.Println()
}
