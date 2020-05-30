//NO default fall through in Go. any switch case condition comes true, it exits the switch condition

package main

import "fmt"

func main() {
	switch "Bond" {
	case "nope":
		fmt.Println("not really")
	case "Bond":
		fmt.Println("the name is Bond, James Bond")
	default:
		fmt.Println("this is the default case")

	}
}
