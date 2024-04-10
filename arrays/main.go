package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in Golang!")

	var b = [3]int{1, 23, 2}
	fmt.Printf("The memory address of b is:%p\n", &b)

	// Assign the array to another variable
	var a = b
	a[0] = 80

	println("The value of a is:", a)
	a[2] = 4
	println(a)

	// Print the memory address of the copied array
	fmt.Println("The a  of a is: \n", &b)
	fmt.Println("The   b is: \n", &a)
	names := [5]string{"abhi", "raj", "ken", "seb"}
	fmt.Println(names, len(names))

	// Slices in go

	var a []int         //un initialized slices
	var b = []int{1, 2} //initalized
	// a = append(a, 1)
	b = append(b, 3)
	b = append(b, 4)
	fmt.Printf("the add of b is %p\n",  &b)
	b = append(b, 5)     
	fmt.Printf("The memory address of b is:%p\n", &b)

	// Assign the array to another variable
	var a = b
	a[0] = 80

	println("The value of a is:", a)
	a[2] = 4
	println(a)

	// Print the memory address of the copied array
	fmt.Println("The a  of a is: \n", &b)
	fmt.Println("The   b is: \n", &a)
	names := [5]string{"abhi", "raj", "ken", "seb"}
	fmt.Println(names, len(names))

	// Slices in go

	var a []int;        
	fmt.Println(b, cap(b))
	fmt.Println(a)
	fmt.Printf("the add of a is %p\n", &a)
	a = b
	d := make([]int, 5)
	fmt.Printf("the type of d is %T\n", d)
	fmt.Println(d)
	e := a
	fmt.Printf("the add of e is %p\n", &e)
	fmt.Println(e)
	e[0] = b[0]
	fmt.Println(e)
	fmt.Println(fmt.Printf("the add of a is: %p", &a)) // this loc is in go's runtime and not actual mem loc
	fmt.Println(b)

}

// Example of go slice pass by ref and value

func doSome(a [3]int, b []int) []int {
	a[0] = 4
	b[0] = 3

	c := make([]int, 5)
	c[4] = 43
	copy(c, b)

	return c
}

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
