package main

import (
	"io"
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("hello world")
	p := make([]byte, 4)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		// error
		// fmt.Println("p:", string(p), "n:", n)
		// true
		fmt.Println("p: ", string(p[:n]))
	}
}
