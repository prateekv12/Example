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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	//fmt.Println(p2)

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlavours {
			fmt.Println(i, val)

		}
	}
}
