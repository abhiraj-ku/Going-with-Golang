package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var name string
	var userRating string
	// Frontend
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your full name!")
	name, _ = reader.ReadString('\n') //the readstring return two things one is

	fmt.Println("Please rate on a scale of 1 to 10!")
	userRating, _ = reader.ReadString('\n')
	mynumRating, _ := strconv.ParseFloat(strings.TrimSpace(userRating), 64)

	// Backend
	fmt.Printf("%v, %v\n", name, mynumRating)

}
