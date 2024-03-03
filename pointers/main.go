package main

import "fmt"

func main() {
	fmt.Println("hello from poimters")

	// var one int = 2
	// var ptr *int = &one
	one := 23
	ptr := &one
	fmt.Println("he actual mem address is ", ptr)
	fmt.Println("the actual values", *ptr)
}
