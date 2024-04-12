package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// a, b := 1, 23
	// c, d := 1.23, 1.45

	// fmt.Printf("%#x %#x\n", a, b)  //hexa-decimal
	// fmt.Printf("%4d %4d\n", a, b) // decimal with four place digit number
	// fmt.Printf("%d %d\n", a, b)   // decimal
	// fmt.Printf("%f %f\n", c, d)   //float

	// s := "a string"

	// b := []byte(s)
	// fmt.Printf("%T\n", s)
	// fmt.Printf("%q\n", s)
	// fmt.Printf("%v\n", s)
	// fmt.Printf("%v\n", b)         // it will return bytes of slice
	// fmt.Printf("%v\n", string(b)) // convert it back to string

	// FILE I/O
	/* package os provides i/o facility
	buffio provides buffered io scanners
	package io/ioutil has extra ability to read and write an entire file  all at once
	package strconv  has utils to convert to/from to string
	*/

	// in every fn return something and err associated with it
	/* in Go EOF is not considered an error
		Copy copies from src to dst until either EOF is reached on src or an error occurs.
		 It returns the number of bytes copied and the first error encountered while copying, if any.

	A successful Copy returns err == nil, not err == EOF.
	 Because Copy is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
	*/
	for _, fname := range os.Args[1:] { // taking input from cmd
		file, err := os.Open(fname) // opening the respected file passed via cmd

		if err != nil { // handling err while file handling
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if _, err := io.Copy(os.Stdout, file); err != nil { // copying ther file to output
			fmt.Fprint(os.Stderr, err)
			continue
		}
		file.Close()
	}

	// ReadAll(file) meaning reading the passed in filename all at once
}
