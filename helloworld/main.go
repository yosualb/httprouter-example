package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index handle request from route /
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	router := httprouter.New()

	router.GET("/", Index)

	log.Fatal(http.ListenAndServe(":8000", router))
}
