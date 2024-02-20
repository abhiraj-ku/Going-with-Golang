// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// fmt.Println("hello world from Go!")
// 	// a := 5
// 	// var bc string
// 	// fmt.Println(bc)
// 	// var ac int
// 	// var acc int = 45
// 	// fmt.Println(ac)
// 	// fmt.Println(acc)

// 	var pi = 3
// 	pi = 6
// 	fmt.Printf("value of pi is %d and its type is %[1]T", pi)
// 	// b := 8.66

// 	// fmt.Printf("the value b is %v and its type is %T\n", b, b)

// 	// var a int = 4
// 	// a = 3
// 	// the below [1] is used to tell %T specifier to use the first params only so we dont have to pass the same param
// 	// for each specifier
// 	// fmt.Printf("the value a is %v and its type is %[1]T\n", a)
// 	// fmt.Println(quote.Go())

// }
package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!") // Sends the response "Hello, World!" to the client
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the About page.") // Sends the response "This is the About page." to the client
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You can contact us at example@example.com.") // Sends the response "You can contact us at example@example.com."
}

func main() {
	http.HandleFunc("/", rootHandler)           // Registers the root handler
	http.HandleFunc("/about", aboutHandler)     // Registers the /about route handler
	http.HandleFunc("/contact", contactHandler) // Registers the /contact route handler

	fmt.Println("Server is listening on port 8000...")
	http.ListenAndServe(":8000", nil) // Starts the server on port 8080
}
