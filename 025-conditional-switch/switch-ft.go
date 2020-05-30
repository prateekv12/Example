//NO default fall through in Go. any switch case condition comes true, it exits the switch condition

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("should not print")
	case (2 == 4):
		fmt.Println("should not print")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough
	case (4 == 4):
		fmt.Println("true, and it prints as it has FT")
	case (5 == 7):
		fmt.Println("should not print")
	}
}
