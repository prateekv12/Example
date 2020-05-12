package main

import "fmt"

func main() {
	fmt.Println("lets go! Vamos!")
	foo()
	fmt.Println("printing something after foo")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I am in foo")
}

//control flow:
//1.sequential
//2. loop / iterative
//3. conditional
