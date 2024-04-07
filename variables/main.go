package main

import "fmt"

// We cannot use := outside or globally
// for that either use var a int =34

// vars with capital name are considered packages and they need to be imported
func main() {

	var name string = "Abhishek kumar"
	fmt.Println(name)
	fmt.Printf("The type of var name is : %T and length is %v \n", name, len(name))

	// numbers
	var numone int8 = 34
	fmt.Print(numone)

	// fmt.Println("hello from vars")
}
