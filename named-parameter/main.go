package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Person handle request from route /person/:name
func Person(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Hello %v!\n", ps.ByName("name"))
}

func main() {
	router := httprouter.New()

	router.GET("/person/:name", Person)

	log.Fatal(http.ListenAndServe(":8000", router))
}
