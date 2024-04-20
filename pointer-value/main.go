package main

import (
	"fmt"
)

// Semantic consistency
// it is that if something is to be passed then always pass it as a pointer
// in go creating something with new doesn't guarantee a heap alloca it can be a stack or heap
func main() {
	fmt.Println("Hello, world!")

	// for loop
	// for i, thing = range things {
	// 	// here the thing is a copy
	// }

	// for i = range things {
	// 	things[i].which = whatever // this is better if we want to change while iterating
	// }

	// Slices
	/*
				anytime a function mutates a slice remember to return the slice afterwards
		because the underlying array could be  reallocated to grow
		if we don't return all the changes could be lost

		we should always keep ref of slices in a loop to expect the desired res ;
		like a local var inside the loop body
	*/

}
