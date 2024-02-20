package main

import (
	"fmt"
)

func main() {

	// str := "Hello from Golang the baap of backend"
	// here str holds data(Helloé) and length of the the string , it is descriptor
	// str := "abhishek"
	// str += " lawde"
	// // str2 := "élite"
	// str2 := str[:4] + "e" + str[:7]

	// str3 := str[:5] + "e" + str[7:]

	// for _, char := range str {
	// 	fmt.Printf("%c ", char)
	// }

	// Split the string into words
	// words := strings.Fields(str)

	// Join the words with a hyphen
	// joined := strings.Join(words, "-")

	// fmt.Println(joined)
	// strings in go is immutable
	// fmt.Printf("%T %[1]v\n", str)
	// fmt.Printf("%8T %[1]v\n", []rune(str))
	// fmt.Printf("%8T %[1]v\n", []byte(str))

	// len(str) is no of bytes req to represent unicode chars and not the unicode chars(in case of unicode )
	// fmt.Println(str, "ka length =", len(str))
	// fmt.Println(str2)
	// fmt.Println(str3)
	// fmt.Println(strings.Contains(str, "abhishe"))

	var a string = "abhishek"
	// var b, _ = fmt.Println(a)
	fmt.Printf("the type of b is %T and its length is %[1]v\n", len(a))
}
