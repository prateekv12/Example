package main

import "fmt"

func main() {
	s := struct {
		first      string
		friends    map[string]int
		favFlavors []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"Z":          888,
		},
		favFlavors: []string{
			"Chocolate",
			"martini",
		},
	}

	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favFlavors)

	for i,v := range s.friends {
		fmt.Println(i,v)
	}
	fmt.Println ("---------------------")
	for k,val := range s.favFlavors {
		fmt.Println(k,val)
	}
}
