package main

import (
	"fmt"
	"time"
)

func generator(limit int, ch chan<- int) {
	for i := 2; i < limit; i++ {
		ch <- i // writes to channel in this loop
	}
	close(ch)
}

// filter function to filter prime number coming to this channel
func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
	close(dst)

}

// sieve function -> basically how long the output should be (limit)
func sieve(limit int) {
	ch := make(chan int)

	go generator(limit, ch)

	// infinite loop to read and filter the result
	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		ch1 := make(chan int)

		go filter(ch, ch1, prime)

		ch = ch1
		fmt.Print(prime, " ")
	}
	fmt.Println()

}

func main() {
	start := time.Now()

	sieve(10000)
	t := time.Since(start).Round(time.Millisecond)

	fmt.Println(t)
}
