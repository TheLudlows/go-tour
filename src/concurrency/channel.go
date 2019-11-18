package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	go func() {
		ch <- 10
		ch <- 20

		fmt.Println("10 put")
	}()
	a := <-ch
	fmt.Println("a", a)
}
