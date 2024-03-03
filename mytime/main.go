package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// formatting the time

	fmt.Println(presentTime.Format("01-02-2006"))
}
