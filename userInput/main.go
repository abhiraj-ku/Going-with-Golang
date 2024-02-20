package main

import (
	"bufio"
	"fmt"
	"os"
)

//	func main() {
//		welcome := "Welcome to user input session!"
//		fmt.Println(welcome)
//		const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
//		// fmt.Println(sample)
//		for i := 0; i < len(sample); i++ {
//			fmt.Printf("%x ", sample[i])
//		}
//	}
// func main() {
// 	var input string
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanln(&input)
// 	fmt.Println("Hello,", input)
// }

func main() {
	fmt.Println("Welcome to input output in golang")

	// create a reader for taking input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Give us a rating")
	input, _ := reader.ReadString('\n')
	fmt.Printf("You have rated us %v star!\n", input)
	fmt.Printf("the type of your input is %T\n", input)
}
