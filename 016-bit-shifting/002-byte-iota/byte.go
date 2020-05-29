package main

import "fmt"

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	// for every step we see a bit shift of 10 bits
	//1024                    10000000000
	//1048576                 100000000000000000000
	//1073741824              1000000000000000000000000000000

}
