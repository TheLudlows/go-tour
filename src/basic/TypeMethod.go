package main

import "fmt"

type Integer int

func (a Integer) compare(b Integer) bool {

	return a > b
}

func main() {
	var a Integer = 10
	var b Integer = 20

	fmt.Println(a.compare(b))
}
