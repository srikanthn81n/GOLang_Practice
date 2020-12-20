package main

import "fmt"

//Passing input through stdin
func main () {
	var min, max int
	fmt.Print("Enter two numbers separated by space:\n")
	fmt.Scanf("%d %d", &min, &max)
	fmt.Println(min, max)


}
