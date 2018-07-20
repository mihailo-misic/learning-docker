package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

func main() {
	// Get all
	router.GET("/", Index)
	// Get one
	// Create
	// Update
	// Delete

	log.Fatal(http.ListenAndServe(":8000", router))
}
