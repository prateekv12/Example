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

	fmt.Println(m["Tony"])

	//there is a way to check if the key exists in the map using the ",ok" idiom

	v, ok := m["Tony"]
	fmt.Println(v)  //will be zero if not present in the map
	fmt.Println(ok) //will be false if its not present in the map

	// using the ",ok" idiom in a if statement

	if v, ok := m["James"]; ok {
		fmt.Println("this is printing from the if statement", v)
	}

}
