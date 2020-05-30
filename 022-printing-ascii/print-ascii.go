//decimal 33 to decimal 122 can be used to print the  ascii characters and the codepoints
//ref: https://en.wikipedia.org/wiki/ASCII

//• decimal value with %d
//• hex value with %#x
//• unicode code point with %#U
//• a tab with \t
//a new line with \n

package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%b\t%#x\t%#U\n", i,i,i,i)
	}

}