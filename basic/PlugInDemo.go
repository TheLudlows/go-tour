package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	filename := "MyPlugin.so"
	p, err := plugin.Open(filename)
	if err != nil {
		log.Fatalf("cannot load plugin %v", filename)
	}
	helloFunc, err := p.Lookup("SayHello")
	if err != nil {
		log.Fatalf("cannot find SayHello in %v", filename)
	}
	helloF := helloFunc.(func(string) string)
	fmt.Println(helloF("aaa"))
}
