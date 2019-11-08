package main

import (
	"fmt"
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
}
