package main

import "fmt"

func main() {

	//the type  map[key]value is a composite literal
	//maps are in the form of key and value

	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	//printing the map
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println(v)
		fmt.Println(ok)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)
}
