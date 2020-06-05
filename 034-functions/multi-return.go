// func (r receiver) identifier(parameters) (return(s)) { ... }
//we define our func with PARAMETERS (if any)
// we call our func and pass in ARGUMENTS (if any)

package main

import "fmt"

func main() {
	foo()
	bar("Rocky")
	s1 := woo("James")
	fmt.Println(s1)

	x, y := mouse("Tony", "Stark")
	fmt.Println(x)
	fmt.Println(y)
}
func foo() {
	fmt.Println("hello from foo")
}

func bar(s string) {
	fmt.Println("Hello,", s)
}

func woo(st string) string {
	return fmt.Sprint("hello from Woo, ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says "Hello"`)
	b := false
	return a, b
}
