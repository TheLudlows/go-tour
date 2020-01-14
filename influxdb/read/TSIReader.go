package main

import (
	"fmt"
	"go-tour/influxdb/utils"
	"io/ioutil"
)

func main() {
	file := "/Users/liuchao56/1"
	data, _ := ioutil.ReadFile(file)

	offBytes := data[len(data)-82:]

	fmt.Println(offBytes)
	fmt.Println(utils.BytesToInt64(offBytes[:8]))
	fmt.Println(utils.BytesToInt64(offBytes[8:16]))
	fmt.Println(utils.BytesToInt64(offBytes[16:24]))
	fmt.Println(utils.BytesToInt64(offBytes[24:32]))
	fmt.Println(utils.BytesToInt64(offBytes[32:40]))
	fmt.Println(utils.BytesToInt64(offBytes[40:48]))
	fmt.Println(utils.BytesToInt64(offBytes[48:56]))
	fmt.Println(utils.BytesToInt64(offBytes[56:64]))
	fmt.Println(utils.BytesToInt64(offBytes[64:72]))
	fmt.Println(utils.BytesToInt64(offBytes[72:80]))
	fmt.Println(utils.BytesToInt16(offBytes[80:82]))

}
