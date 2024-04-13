package main

import (
	"fmt"
)

// single return functions
// func do_some(a, b int) int {
// 	return a + b
// }

// multiple return functions

// func return_many(a string, b int) (int, string) { // (int,string) tells the output expected striclty in this order
// 	return b, a
// }

// Naked Function in golang

/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function.
A return statement without arguments returns the named return values. This is known as a "naked" return.
*/

// func split(sum int) (x, y int) { // here we are expecting return types as named var so, they are treated as variables defined at the top of the function.
// 	x = sum * 4 / 9
// 	y = sum - x
// 	return
// }

func main() {
	fmt.Println("Hello, world!")

	// fmt.Println(do_some(2, 4))
	// a, b := return_many("hello", 8)
	// fmt.Println(a, b)
	// fmt.Println(split(17))

	// pass by value
	/*
		numbers, bool , arrays , structs

	*/
	// pass by reference
	/*
		pointers, strings, slices, maps ,channels

	*/

	// defer in go :defer in Go is used for deferring function calls until the surrounding function returns.

	a := 10
	defer fmt.Println(a)

	a = 12
	fmt.Println(a)
}
