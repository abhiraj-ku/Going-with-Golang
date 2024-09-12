package main

import (
	"log"
	"net/http"
	"time"
)

// this type will hold the url,err,and time it takes to recieve the data (latency)
type result struct {
	url     string
	err     error
	latency time.Duration
}

// make the call to get the data
func get(url string, ch chan<- result) { // takes a channel which can only writes to
	start := time.Now()
	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}

	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t} // channel which reads only
		resp.Body.Close()
	}
}

func main() {
	result := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://google.com",
		"https://wsj.com",
		"https://abhikumar.site",
	}

	for _, url := range list {
		go get(url, result)
	}
	for range list {
		r := <-result

		if r.err != nil {
			log.Printf("%-20s %s\n", r.url, r.err)
		} else {
			log.Printf("%-20s %s\n", r.url, r.latency)

		}

	}

}
