//print even numbers between 1 and 100 using the break continue logic

package main

import "fmt"

func main() {

	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}
}
