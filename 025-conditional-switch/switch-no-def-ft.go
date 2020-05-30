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
	case (4 == 4):
		fmt.Println("true, but i dont think it prints") // does not print even though its true because there is no default fall thru
	case (5 == 7):
		fmt.Println("should not print")
	}
}
