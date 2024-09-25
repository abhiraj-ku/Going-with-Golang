package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func readFile(filePath string, wg *sync.WaitGroup, lines chan<- string) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file", err)

	}
}

func main() {
	start := time.Now()

	fmt.Println("Welcome to paralle log processor!")

	logFiles := []string{"access1.txt", "access2.txt", "access3.txt", "access4.txt", "access5.txt", "access6.txt", "access7.txt", "access8.txt"}
	lines := make(chan string, 100)
	var wg sync.WaitGroup

	for _, filePath := range logFiles {
		wg.Add(1)
		go readFile(filePath, &wg, lines)
	}

	go func() {
		wg.Wait()
		close(lines)

	}() // close when all goroutines are done

	for line := range lines {
		fmt.Println(line)
	}
	t := time.Since(start)
	fmt.Print(start)
	fmt.Print(t)
}
