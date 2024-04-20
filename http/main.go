package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	for i := 0; i < 1000; i++ {

		fmt.Fprintf(w, "Hello world from :%s\n", r.URL.Path[:])
	}
	duration := time.Since(start)
	fmt.Printf("time taken for response:%v\n", duration)
}
func main() {
	fmt.Println("Hello, http")

	http.HandleFunc("/", homePageHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
