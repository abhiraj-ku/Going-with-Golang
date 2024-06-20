package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Hello, world!")
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Environment variable PORT is not set")
	} else {
		fmt.Println("Port:", port)
	}
}
