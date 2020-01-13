package main

import "fmt"

type people People

func main() {
	fmt.Println("hello word")
	variableDefinition()
	address()
	iotaFunc()
	a, b := muxReturn()
	fmt.Println(a, b)
	addFunc := func(a int, b int) int {
		return a + b
	}
	println(addFunc(999, 1))
	p := people{100, "four"}
	pAddress := &p
	eq := &(*pAddress) == pAddress
	fmt.Println(p.name, eq)
}

func variableDefinition() {
	var age = 10
	var name = "four"
	height := 180.12
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(height)
}

func address() {
	var age int = 10
	var address = &age
	fmt.Println(age)
	*address = 100
	fmt.Println(address)
	fmt.Println(age)

}

func constFunc() {
	const AGE = 100
	fmt.Println(AGE)
}

func iotaFunc() {
	const (
		a = iota
		b = iota
	)
	fmt.Println(a, b)
}
func muxReturn() (string, string) {
	return "google", "IBM"
}

type People struct {
	age  int
	name string
}
