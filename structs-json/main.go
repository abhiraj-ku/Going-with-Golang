package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}
type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println("Hello, world!")

	fmt.Println("Hello Structs!")
	a := Animal{"cat", 19}
	b, _ := json.Marshal(a)

	fmt.Printf("%s\n", b)
	fmt.Printf("Type = %T, value =%+[1]v\n", a)
	// var e Employee
	// e.Name = "Abhishek kumar"
	// e.Number = 1234
	// // Initialize boss
	// boss := Employee{Name: "Boss Name", Number: 9999} // Example boss data
	// e.Boss = &boss

	// fmt.Printf("%T %+[1]v", e)
	// if e.Boss != nil {
	// 	fmt.Println("Boss Data:")
	// 	fmt.Printf("Name: %s\nNumber: %d\n", e.Boss.Name, e.Boss.Number)
	// } else {
	// 	fmt.Println("\nNo boss assigned")
	// }

	// using the struct literal

	// var a = Employee{
	// 	"matt",
	// 	1,
	// 	nil,
	// 	time.Now(),
	// }

	// the above way has a problem cause now we have to give values for each of the field in the struct

	// to avoid the above scenario we need to pass the values as in the original struct
	// var b = Employee{
	// 	Name:   "matt",
	// 	Number: 1,
	// 	Hired:  time.Now(),
	// }
	// fmt.Printf("%T %+[1]v", b)

	// Another way is to make a annonymous struct OR a Struct literal;
	var g = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The white album",
		"the beatles",
		23,
		2323,
	}
	fmt.Println(g)

	var palbum = &g
	fmt.Println(palbum)

	// Struct compatibility
	/* structs are copatible if and only if
	1. the fields have same type and name
	2. in the same order
	3.and the same tags(*)

	Note: a strut may be copied or passed as a param , the zero value of struct is "zero " for all its fields

	so copatibilty is the issue here for the structs data types as long as they have the same value its okay
	*/

	// structs are passed as a copy to the function  params unless a pointer is used

	// EMpty structs: it us useful as it has no memory

	// var isPresent map[int]struct{}
	// done := make(chan struct{})

	// struct are passed by value;
	// struct with no field is useful because thi does not takes up space
}
