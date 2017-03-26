package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func lazyHandler(w http.ResponseWriter, r *http.Request) {
	randNum := rand.Intn(2)
	if randNum == 0 {
		time.Sleep(6 * time.Second)
		fmt.Fprintf(w, "slow response, %d\n", randNum)
		fmt.Printf("slow response, %d\n", randNum)
		return
	}

	fmt.Fprintf(w, "fast response, %d\n", randNum)
	fmt.Printf("fast response, %d\n", randNum)
	return
}

func main() {
	http.HandleFunc("/lazy", lazyHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
