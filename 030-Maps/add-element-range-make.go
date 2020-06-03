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

	//printing the value using the key

	fmt.Println(m["James"])

	//if you try to print a value of a non-existent key then it prints 0.

	// using the ",ok" idiom in a if statement

	if v, ok := m["James"]; ok {
		fmt.Println("this is printing from the if statement", v)
	}

	// adding an element to the map

	m["Todd"] = 33

	// for loop to print key and values in a map

	for k, v := range m {
		fmt.Println(k, v)
	}

	// for loop to print index and values in a key - detour example

	xi := []int{2, 3, 4, 5, 6}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}
