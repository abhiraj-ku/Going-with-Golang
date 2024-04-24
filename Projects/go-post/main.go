package main

import (
	"example/go-post/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Everything all correct")

	r := router.Router()
	fmt.Println("Starting server on port 7070")
	log.Fatal(http.ListenAndServe(":7070", r))

}
