package main

import "fmt"

type my_var int

var x my_var

func main() {

	//x = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
