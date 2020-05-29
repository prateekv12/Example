package main

import "fmt"

const (
	_  = iota             //discarding the first iota = 0
	kb = 1 << (iota * 10) // iota = 1
	mb = 1 << (iota * 10) //iota = 2
	gb = 1 << (iota * 10) // iota = 3
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}

//OUTPUT

//$ go run iota.go
//1024                    10000000000
//1048576                 100000000000000000000
//1073741824              1000000000000000000000000000000
