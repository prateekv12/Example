package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[4] = 24
	fmt.Println(x)
	fmt.Println(len(x))

}
