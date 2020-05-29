package main

import "fmt"

//const a = 42
//const b = 45.56
//const c = "Yowza folks"

// constants can be declared together without using cosnt a million times.

//the below is called untyped constants
const (
	a = 42
	b = 45.56
	c = "Yowza folks"
)

//the below is called typed constants (less flexible)
//const (
//	a int     = 42
//	b float64 = 45.56
//	c string  = "Yowza folks"
//)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
