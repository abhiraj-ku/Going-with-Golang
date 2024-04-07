package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays in Golang!")

	var b = [3]int{1, 23, 2}
	fmt.Printf("The memory address of b is:%p\n", &b)

	Assign the array to another variable
	var a = b
	a[0] = 80

	println("The value of a is:", a)
	a[2] = 4
	println(a)

	Print the memory address of the copied array
	fmt.Println("The a  of a is: \n", &b)
	fmt.Println("The   b is: \n", &a)
	names := [5]string{"abhi", "raj", "ken", "seb"}
	fmt.Println(names, len(names))

	Slices in go

	var a []int         //un initialized slices
	var b = []int{1, 2} //initalized
	// a = append(a, 1)
	b = append(b, 3)
	b = append(b, 4)
	fmt.Printf("the add of b is %p\n", &b)
	b = append(b, 5) //from here on it creates a new array with larger cap and copies to the prev array
	fmt.Printf("the add of b is %p\n", &b)
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
