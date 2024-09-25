package main

import (
	"log"
	"time"
)

// func main() {
// 	c1 := make(chan int)
// 	c2 := make(chan string)

// 	go func() {

// 		c1 <- 42
// 	}()

// 	go func() {
// 		c2 <- "Hello"
// 	}()

// 	for i := 0; i < 12; i++ {
// 		time.Sleep(time.Duration(1) * time.Second)
// 		v1 := <-c1
// 		log.Println("Received from c1:", v1)
// 		v2 := <-c2
// 		log.Println("Received from c2:", v2)
// 		// select {
// 		// case v1 := <-c1:
// 		// 	log.Println("Received from c1:", v1)
// 		// case v2 := <-c2:
// 		// 	log.Println("Received from c2:", v2)
// 		// default:
// 		// 	log.Println("Received from none")
// 		// }

//		}
//	}
// func main() {
// 	// ch := make(chan int)

// 	// go func() {
// 	// 	log.Println("Goroutine: About to send")
// 	// 	ch <- 42
// 	// 	log.Println("Goroutine: Sent")
// 	// }()

// 	// time.Sleep(2 * time.Second)
// 	// log.Println("Main: About to receive")
// 	// value := <-ch
// 	// log.Println("Main: Received", value)
// 	ch := make(chan int)
// 	timeout := time.After(3 * time.Second)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		log.Println("Goroutine: About to send")

// 		ch <- 42
// 		log.Println("Goroutine: Sent")
// 	}()

// 	select {
// 	case value := <-ch:
// 		log.Println("Main: Received", value)
// 	case <-timeout:
// 		log.Println("Main: Timed out")
// 	}

// 	// Try to receive again, this time with a default case
// 	select {
// 	case value := <-ch:
// 		log.Println("Main: Received another value", value)
// 	default:
// 		log.Println("Main: No value available, not blocking")
// 	}

// 	log.Println("Main: Exiting")
// }

// MATT HOLIDAY CODE
func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}

		}(i+1, chans[i])
	}

	for i := 0; i < 12; i++ {
		// t1 := <-chans[0] // polling of result means waiting for first one to complete
		// t2 := <-chans[1]

		// log.Printf("recieved from %d:", t1)
		// log.Printf("recieved from %d:", t2)

		select {
		case t1 := <-chans[0]:
			log.Printf("recieved from %d", t1)
		case t2 := <-chans[1]:
			log.Printf("recieved from %d", t2)

		}
	}

}
