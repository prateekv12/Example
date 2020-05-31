package main

import "fmt"

func main() {
	x := "James Bond"

	if x == "Dr No" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println(x)
	} else {
		fmt.Println("x is neither Dr No or James Bond")
	}
}
