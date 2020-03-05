package main

import (
	"fmt"
	"reflect"
)

type Printer interface {
	print(str string)
}

type MyPrinter struct {
}

func (MyPrinter) print(str string) {
	fmt.Println("hi", str)
}

func main() {

	var printer Printer
	printer = new(MyPrinter)
	p := printer.print
	p("you")
	printer.print("four")
	var v1 interface{} = 1
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(&v1)

	type User struct {
		id   int
		name string
	}

	u := User{1, "Tom"}
	var i interface{} = u
	u.id = 2
	u.name = "Jack"
	fmt.Println(i.(User).id)
	fmt.Printf("%v\n", u)
	fmt.Printf("%v\n", i.(User))

}
