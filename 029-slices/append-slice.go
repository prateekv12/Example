package main

import "fmt"

func main() {
	//COMPOSITE LITERAL

	//x := type{values}

	x := []int{11, 22, 33, 44, 55, 66}
	fmt.Println(x)
	x = append(x, 15, 16, 17, 19)
	fmt.Println(x)
	y := []int{111, 222, 333, 444, 555}
	x = append(x, y...)
	fmt.Println(x)
}
