package main

import (
	"io"
	"log"
	"net/http"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

type MyHandle struct{}

func (m *MyHandle) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "URL:"+req.URL.String())
	if h, ok := mux[req.URL.String()]; ok {
		h(w, req)
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello")
}

func bye(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bye")
}

func main() {
	//mux := http.NewServeMux()
	//mux.Handle("/", &MyHandle{})

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = hello
	mux["/bye"] = bye

	server := http.Server{
		Addr: ":9999",
		Handler: &MyHandle{},
	}
	//http.ListenAndServe(":9999", mux)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("http listenAndServe error")
	}
}
