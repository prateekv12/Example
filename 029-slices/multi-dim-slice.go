package main

import "fmt"

func main() {
	//COMPOSITE LITERAL

	//x := type{values}

	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	//creating a slice of a slice ==> MULTIDIMENSIONAL SLICE

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	//prints the multi-dimensional slice
}
