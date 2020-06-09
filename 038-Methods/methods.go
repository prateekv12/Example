// When you attach a func to a type it is a method of that type. You attach a func to a type with a RECEIVER.

package main

import "fmt"

type person struct {
	first string
	last  string
}

type SecretAgent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }
//attaching a function to any value of type "SecretAgent"

func (s SecretAgent) speak() {
	fmt.Println("I am ", s.first, s.last)
}
func main() {

	sa1 := SecretAgent{
		person: person{
			"James",
			"bond",
		},
		ltk: true,
	}

	sa2 := SecretAgent{
		person: person{
			"Tony",
			"Stark",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

}
