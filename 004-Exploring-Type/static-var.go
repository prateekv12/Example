package main

import (
	"fmt"
)

var y = 42

//Declare a variable with identifier z
//of Type String
//Assign the value "Shaken, not Stirred"

// STATIC PROGRAMMING LANGUAGE

var z = "Shaken, not Stirred"

var a string = `James Said "Shaken, not Stirred" `

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
