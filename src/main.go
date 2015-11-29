package main

import (
	"log"
	"net/http"
	"router"
)

func main() {

	mainRouter := router.Router()

	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
