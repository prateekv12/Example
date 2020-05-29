package main

import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	//TO DEMONSTRATE BIT SHIFTING LETS ASSIGN ANOTHER VARIABLE y and shift bits by 1.

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

}
