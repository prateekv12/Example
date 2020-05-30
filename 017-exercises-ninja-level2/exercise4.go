package main

import "fmt"

//var x int

func main() {
	x := 24
	fmt.Printf("%d\t%b\t%#x\t\n", x, x, x)

	y := x << 1
	fmt.Printf("%d\t%b\t%#x\t", y, y, y)

}
