package main

import "fmt"

func main() {

	birthyear := 1990
	for {
		fmt.Println(birthyear)
		birthyear++
		if birthyear > 2020 {
			break

		}

	}
}
