package main

import (
	"fmt"
	"github.com/jwilder/encoding/simple8b"
)

func main() {

	arr := make([]uint64, 240)
	fmt.Println(len(arr))
	des, err := simple8b.EncodeAll(arr)
	fmt.Println(des, err)
	deco := make([]uint64, 400)
	simple8b.DecodeAll(deco, des)
	fmt.Println(len(deco))

}
