//NO default fall through in Go. any switch case condition comes true, it exits the switch condition

package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "nope", "Bond", "Dr No":
		fmt.Println("This is nope or Bond or Dr No")
	case "M":
		fmt.Println("the name is M")
	default:
		fmt.Println("this is the default case")

	}
}
