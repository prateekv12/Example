package main

import "fmt"

type person struct {
	first       string
	last        string
	favFlavours []string
}

func main() {

	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavours: []string{
			"Chocolate",
			"Martini",
			"Mango",
		},
	}

	//fmt.Println(p1)

	p2 := person{
		first: "Tony",
		last:  "Stark",
		favFlavours: []string{
			"Vanilla",
			"Cookie Dough",
			"Coffee",
		},
	}

	//fmt.Println(p2)

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavours {
		fmt.Println(i, v)
	}
}
