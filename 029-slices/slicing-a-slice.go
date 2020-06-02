package main

import "fmt"

func main() {
	//COMPOSITE LITERAL

	//x := type{values}

	x := []int{11, 22, 33, 44, 55, 66}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}
}
