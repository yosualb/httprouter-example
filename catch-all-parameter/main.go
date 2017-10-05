package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Say handle request from route /say/*everything
func Say(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Remove front /
	everything := ps.ByName("everything")[1:]
	fmt.Fprintf(w, "Say: %v\n", everything)
}

func main() {
	router := httprouter.New()

	router.GET("/say/*everything", Say)

	log.Fatal(http.ListenAndServe(":8000", router))
}
