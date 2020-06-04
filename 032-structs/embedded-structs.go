//Structs is a composite data type or aggregate data structure

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type SecretAgent struct {
	person
	ltk bool
}

func main() {

	sa1 := SecretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   44,
		},
		ltk: true,
	}

	sa2 := SecretAgent{
		person: person{
			first: "Tony",
			last:  "Stark",
			age:   37,
		},
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)
	fmt.Println(sa2.first, sa2.last, sa2.age)

}
