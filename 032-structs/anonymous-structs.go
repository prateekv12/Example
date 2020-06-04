//called anonymous as it does not have a predefined name and is used as a composite literal type while declaring the struct
//used when you are trying to avoid code pollution
//Structs is a composite data type or aggregate data structure

package main

import "fmt"

func main() {

	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   42,
	}

	p2 := struct {
		first string
		last  string
		age   int
	}{
		first: "Tony",
		last:  "Stark",
		age:   37,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

}
