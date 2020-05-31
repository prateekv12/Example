package main

import "fmt"

func main() {

	for i := 10; i <= 100; i++ {
		j := (i % 4)
		fmt.Printf("when %v is divided by 4, the reminder is %v\n", i, j)
	}
}
