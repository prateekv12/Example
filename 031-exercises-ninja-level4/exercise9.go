package main

import "fmt"

func main() {

	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`Stark_Tony`] = []string{`Being Cool`, `Shawrma`, `gadgets`}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println("this is the value for the record", k)

		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
