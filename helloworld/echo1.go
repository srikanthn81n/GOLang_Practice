package main

import (
	"fmt"
	"os"
)
//Accepting command line arguments with os.Args module and printing arguments
//steps to perform: go build echo1.go
//./echo1.exe arg1 arg2
//Here in this example for loop with range parameter is utilized
func main () {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
		fmt.Println("element at index:", i, "is:", arg  )
	}
	fmt.Println(s)
}
