package main

import "fmt"

const (
	a = 42
	b = "Why"
)

const (
	c int    = 99
	d string = "Help"
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
}
