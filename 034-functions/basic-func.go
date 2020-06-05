// func (r receiver) identifier(parameters) (return(s)) { ... }
//we define our func with PARAMETERS (if any)
// we call our func and pass in ARGUMENTS (if any)

package main

import "fmt"

func main() {
	foo()
}
func foo() {
	fmt.Println("hello from foo")
}
