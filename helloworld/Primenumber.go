package main

import "fmt"

func main () {

	//for loop to find prime numbers between 1 and 4000
	for i :=1; i <= 4000; i++ {
		primen := true
		for c :=2; c <= i/2 ; c++{
			if i%c == 0 {
				primen = false
				break
			}

			}
		if primen == true && i !=1{
			fmt.Printf("%d\t", i)

		}

	}
}
