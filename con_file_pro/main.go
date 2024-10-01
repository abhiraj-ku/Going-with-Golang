package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

//done <- c // writing to channel done
//c <-done reading from done

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

	bytes, err := io.Copy(hash, file)
	fmt.Printf("file size: %.7f MB\n", float64(bytes)/1048576)
	if err != nil {
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

// function to process files and send their hashes to the pairs channel
func processFiles(paths <-chan string, pairs chan<- pair, done chan<- bool) {
	for path := range paths {
		pairs <- hashFile(path)
	}
	done <- true
}

func walkDir(dir string, paths chan<- string, wg *sync.WaitGroup) error {
	defer wg.Done()

	visit := func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.Mode().IsRegular() && fi.Size() > 0 {
			paths <- p
		}
		return nil
	}
	return filepath.Walk(dir, visit)
}

func main() {
	start := time.Now()
	if len(os.Args) < 2 {
		log.Fatal("please provide a directory to scan")
	}
	dir := os.Args[1]

	// uses 2 cores of cpu
	worker := 2 * runtime.GOMAXPROCS(0)

	paths := make(chan string, worker)
	pairs := make(chan pair, worker)
	done := make(chan bool, worker)
	result := make(chan results)

	// start the hash collector
	go collectHashes(pairs, result)

	// start the goroutines
	for i := 0; i < worker; i++ {
		go processFiles(paths, pairs, done)
	}

	// walk the dir concurently
	var wg sync.WaitGroup

	wg.Add(1)

	if err := walkDir(dir, paths, &wg); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	close(paths)

	// wait for all goroutines to finish
	for i := 0; i < worker; i++ {
		<-done // read true for all routines
	}
	close(pairs)

	// print the results
	hashes := <-result // read the result channel

	for hash, files := range hashes {
		if len(files) > 1 {
			fmt.Printf("Duplicate files for hash : %s\n", hash[len(hash)-7:])
			for _, file := range files {
				fmt.Println(file)
			}
		}
	}

	time := time.Since(start)

	fmt.Println(time)

}
