//NO default fall through in Go. any switch case condition comes true, it exits the switch condition

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("should not print")
	case (2 == 4):
		fmt.Println("should not print")
	default:
		fmt.Println("this is the default case")

	}
}
