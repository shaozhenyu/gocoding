package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

func main() {
	b := bytes.NewReader([]byte("heihei mua!\n"))

	reader := io.LimitReader(b, 5)

	if _, err := io.Copy(os.Stdout, reader); err != nil {
		log.Fatal(err)
	}
}

