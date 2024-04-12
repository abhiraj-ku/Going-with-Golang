package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang!")

	var b = [3]int{1, 23, 2}
	fmt.Printf("The memory address of b is:%p\n", &b)

	// Assign the array to another variable
	var a = b
	a[0] = 80

	fmt.Println("The value of a is:", a)
	a[2] = 4
	fmt.Println(a)

	// Print the memory address of the copied array
	fmt.Println("The address of a is: \n", &b)
	fmt.Println("The address of b is: \n", &a)

	names := [5]string{"abhi", "raj", "ken", "seb"}
	fmt.Println(names, len(names))

	// Slices in Go

	var c []int         // uninitialized slice
	var d = []int{1, 2} // initialized slice
	d = append(d, 3)
	d = append(d, 4)
	fmt.Printf("The memory address of d is:%p\n", &d)
	d = append(d, 5)

	fmt.Println("The value of d is:", d)
	d[2] = 4
	fmt.Println(d)

	// Print the memory address of the copied slice
	fmt.Println("The address of c is: \n", &c)
	fmt.Println("The address of d is: \n", &d)

	// Slices in Go

	var aSlice []int
	fmt.Println(d, cap(d))
	fmt.Println(aSlice)
	fmt.Printf("The address of aSlice is %p\n", &aSlice)
	aSlice = d
	e := make([]int, 5)
	fmt.Printf("The type of e is %T\n", e)
	fmt.Println(e)
	f := aSlice
	fmt.Printf("The address of f is %p\n", &f)
	fmt.Println(f)
	f[0] = d[0]
	fmt.Println(f)
	fmt.Println(fmt.Printf("The address of aSlice is: %p", &aSlice)) // this location is in Go's runtime and not actual memory location
	fmt.Println(d)

}

// Example of Go slice pass by reference and value

// func doSome(a [3]int, b []int) []int {
// 	a[0] = 4
// 	b[0] = 3

// 	c := make([]int, 5)
// 	c[4] = 43
// 	copy(c, b)

// 	return c
// }

// func main() {
// 	var w = [...]int{1, 2, 3} // array of len(3)
// 	var x = []int{0, 0, 0}    // slice of len(3)

// 	y := doSome(w, x)
// 	fmt.Println("Array w:", w)
// 	fmt.Println("Slice x:", x)
// 	fmt.Println("Returned slice y:", y)
// }

// MAPS in Go := it is a dictionary with key and value

// func main() {
// 	var m = map[string]int{
// 		"and": 1,
// 		"the": 1,
// 		"or":  1,
// 	}

// 	d := len(m)
// 	e := m["df"]
// 	fmt.Println(d)
// 	fmt.Println(e)

// }
