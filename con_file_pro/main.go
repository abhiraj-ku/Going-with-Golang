package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

type pair struct {
	hash, path string
}

type results map[string][]string

// hash using md5
func hashFile(path string) pair {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hash := md5.New()

	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
	}
	return pair{fmt.Sprintf("%x", hash.Sum(nil)), path}
}

// collect hashes and store it into results channel
func collectHashes(pairs <-chan pair, result chan<- results) {
	hashes := make(results)
	for p := range pairs {
		hashes[p.hash] = append(hashes[p.hash], p.path)
	}
	result <- hashes
}
