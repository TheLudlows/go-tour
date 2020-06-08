package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "aabbcc dd eee"

	for a, b := range str {
		fmt.Println(a, b)
	}

	fmt.Println(unicode.IsLetter(32))
	fmt.Println(unicode.IsLetter('a'))
	fmt.Println(unicode.IsLetter('-'))

}
