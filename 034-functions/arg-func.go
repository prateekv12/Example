// func (r receiver) identifier(parameters) (return(s)) { ... }
//we define our func with PARAMETERS (if any)
// we call our func and pass in ARGUMENTS (if any)

package main

import "fmt"

func main() {
	foo()
	bar("Rocky")
}
func foo() {
	fmt.Println("hello from foo")
}

//Everything in Go is pass by VALUE

func bar(s string) {
	fmt.Println("Hello,", s)
}
