//the keyword iota represents successive integer constants 0,1,2..
//It resets to 0 whenever the word const appears in the source code, and increments after each const specification.

package main

import "fmt"

const (
	a = iota
	b
	c
)

//It resets to 0 whenever the word const appears in the source code, and increments after each const specification.

const (
	d = iota
	e
	f
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
