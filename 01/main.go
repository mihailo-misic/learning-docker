package main

import (
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

func main() {
	router.GET("/", Home)
	
	log.Fatal(http.ListenAndServe(":8000", router))
}
