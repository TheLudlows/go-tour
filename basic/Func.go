package main

import "fmt"

func main() {
	f := func(str string) string {
		fmt.Println(str)
		return str
	}

	f("bbbb")
	panic(f("aaaaa"))

}
