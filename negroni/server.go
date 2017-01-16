package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

type MyHandle int

func (m MyHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/myhandle" {
		fmt.Fprintf(w, "myHandle %d\n", m)
	}
}

func MyMiddleWare1(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(w, "before\n")
	next(w, r)
	fmt.Fprintf(w, "after\n")
}

func MyMiddleWare2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Fprintf(w, "before2\n")
	next(w, r)
	fmt.Fprintf(w, "after2\n")
}

func main() {

	var i MyHandle
	i = 10

	mux := http.NewServeMux()

	newMux := http.NewServeMux()

	mux.HandleFunc("/test1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test1")
	})

	mux.HandleFunc("/test2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test2")
	})

	newMux.HandleFunc("/group", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/group/test")
	})

	mux.Handle("/group", negroni.New(
		negroni.Wrap(newMux),
	))

	mux.Handle("/myhandle", i)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(MyMiddleWare1))
	//n.UseFunc(MyMiddleWare2)
	n.UseHandler(mux)
	n.UseFunc(MyMiddleWare2)
	n.Run(":3000")
}
