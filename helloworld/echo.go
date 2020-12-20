package main

import (
	"fmt"
	"os"
)

//Accepting command line arguments with os.Args module and printing arguments
//steps to perform: go build echo.go
//./echo.exe arg1 arg2
func main () {
	var s, sep string
	argswithcmd := os.Args
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
		//fmt.Println(s)
	}

	fmt.Println(s)
	fmt.Println(argswithcmd)
	fmt.Println(argswithcmd[1:])
}
