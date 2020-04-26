package main

import "fmt"

func main() {
	fmt.Println("main")
	var str string = "a"
	switch str {
	case "a":
		fmt.Println("a")
	default:
		return
	}
	fmt.Println("b")

}
func init() {
	fmt.Println("init")
}
