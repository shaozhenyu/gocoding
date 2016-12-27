package main

import (
	"fmt"
	"io"
	"log"
	"reflect"
	"os"
	"strings"
)

func main() {
	str := strings.NewReader("hello golang\n")
	strT := reflect.TypeOf(str)
	fmt.Println(strT.Kind())
	
	written, err := io.Copy(os.Stdout, str);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(written)
}
