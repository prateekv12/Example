package main

import "fmt"

func main() {

	x := 44

	if x == 40 {
		fmt.Println("the value of x is 40")
	} else if x == 41 {
		fmt.Println("the value of x is 41")
	} else if x == 42 {
		fmt.Println("the value of x is 42")
	} else if x == 43 {
		fmt.Println("the value of x is 43")
	} else {
		fmt.Println("the value of x is NOT 40, 41, 42 OR 43")
	}

}
