package main

import "fmt"

func main() {

	x := []string{"James", "Bond", "Shaken, not stirred"}
	y := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	z := [][]string{x, y}
	fmt.Println(z)

	fmt.Println(len(z))

	for k, xs := range z {
		fmt.Println("record\n", k)
		for i, v := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v\n", i, v)
		}
	}

}
