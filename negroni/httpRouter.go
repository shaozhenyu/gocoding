package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func GetName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "httprouter get:name %s\n", ps.ByName("name"))

}

// match all /file/
func GetFilePath(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "httprouter get *filepath %s\n", ps.ByName("filepath"))
}

func main() {
	router := httprouter.New()
	router.GET("/get/:name", GetName)
	router.GET("/file/*filepath", GetFilePath)

	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}
