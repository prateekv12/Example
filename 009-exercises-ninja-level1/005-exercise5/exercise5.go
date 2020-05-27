package main

import "fmt"

type new_var string

var x new_var
var y string

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = "Bond, James Bond"
	y = "You can't see me!"
	fmt.Println("the value of x is ", x)
	fmt.Println("the value of y is ", y)
	fmt.Printf("%T\n", y)
	//converting the type of var x to var y
	y = string(x)
	fmt.Println("the new value of y is ", y)
	fmt.Printf("%T\n", y)
}
