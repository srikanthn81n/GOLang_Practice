package main

import "fmt"

func main() {
	var i string
	var k string
	fmt.Printf("Message of the day\n")
	fmt.Printf("Enter name of the city:\n")
	fmt.Scanf("%s\n", &i)
	fmt.Printf("Enter name :\n")
	fmt.Scanf("%s\n", &k)
	fmt.Printf("Print the value of city: %s %s", i,k)
}
