package main

import "fmt"

func main() {
	func() {
		fmt.Println("匿名函数执行")
	}()

}
