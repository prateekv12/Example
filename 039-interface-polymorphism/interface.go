//An interface allows a value to be of more than one type. We create an interface using this syntax: “keyword identifier type”
//so for an interface it would be: “type human interface” We then define which method(s) a type must have to implement that interface.

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

//A VALUE CAN BE OF MORE THAN ONE TYPE

type human interface {
	speak()
}

func (s SecretAgent) speak() {
	fmt.Println("I am ", s.first, s.last, "- the SecretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am ", p.first, p.last, "- the person speak")
}

func bar(h human) {
	fmt.Println("I was passed into bar", h)
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

	p1 := person{
		first: "John",
		last:  "Cena",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()
	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)

}
