package main

import "fmt"

func main() {
	s := "Hello, Playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) //strings are slice of bytes. using coding scheme (ASCII) to represent the letter of alphabets
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) //utf 8 (code point ) coding scheme representation of the characters of the string s. each codepoint is called a rune.
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v) //Prints out the decimal values of all characters of the string
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v) // this is to print out the HEX values
	}
}
