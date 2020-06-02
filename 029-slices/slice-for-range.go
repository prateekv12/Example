package main

import "fmt"

func main() {
	//COMPOSITE LITERAL

	//x := type{values}

	x := []int{11, 22, 33, 44, 55, 66}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(len(x))

	//LOOPING THROUGH THE SLICE USING FOR RANGE

	for i, v := range x {
		fmt.Println(i, v)
	}
}

//slices allow you to group the VALUES of same TYPE
