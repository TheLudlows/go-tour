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
	printer.print("four")
	var v1 interface{} = 1
	fmt.Println(reflect.TypeOf(v1))
	fmt.Println(&v1)

}
