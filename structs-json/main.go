package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {
	fmt.Println("Hello, world!")
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
	var b = Employee{
		Name:   "matt",
		Number: 1,
		Hired:  time.Now(),
	}
	fmt.Printf("%T %+[1]v", b)

}
