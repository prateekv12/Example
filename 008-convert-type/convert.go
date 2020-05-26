package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b = 99
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// CONVERTING THE DATA TYPE OF VAR b to VAR a

	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
