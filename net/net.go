package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			resp, err := http.PostForm()
			if err != nil {
				fmt.Println(err)
				return
			}
			resp.Body.Close()
		}()
	}
	fmt.Println("ok")
	time.Sleep(100 * time.Second)
	fmt.Println("ok1")
}
