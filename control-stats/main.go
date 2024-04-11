package main

import "fmt"

func main() {
	// if-else
	// in go it is neccessary to have {brakcets even for single line too}

	// in go we can also do declarations before actual if condition

	//for loop method -> 1
	for i := 0; i <= 9; i++ {
		fmt.Printf("(%d,%d)\n", i, i*i)
	}

	//for loop method -> 2 for condition{}
	sum := 0
	count := 0
	max := 100

	for sum <= max {
		sum += count
		count++
	}
	countplus := count
	count++
	count++
	fmt.Printf("the value of count: %d and after plus %d", countplus, count)

	//for loop method -> 3 index,value {}
	fruits := [5]string{"hello", "baby", "how", "are", "you"}
	fmt.Print(fruits)

	fruitmap := map[int]string{
		20: "hello",
		30: "how",
		40: "are",
		50: "ya",
	}

	for i, v := range fruitmap {
		// the order of element in map in go is random so everytime it will be new output
		fmt.Printf("index:%d value:%s\n", i, v)
	}

	// switch-case statements in golang

}
