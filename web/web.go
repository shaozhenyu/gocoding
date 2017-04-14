package main

import (
	"fmt"
	"log"
	"net/http"
)

type Router struct {
	Name string
}

func (r *Router) ServeHTTP(http.ResponseWriter, *http.Request) {
	fmt.Println("new route")
}

func New() *Router {
	return &Router{
		Name: "aa",
	}
}

func main() {
	fmt.Println("test")

	router := New()


	http.HandleFunc("/hello", Hello0)
	http.HandleFunc("/hello1", Hello1)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Hello1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello1111111")
}

func Hello0(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello00000000")
}
