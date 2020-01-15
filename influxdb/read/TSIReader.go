package main

import (
	"fmt"
	"go-tour/influxdb/utils"
	"io/ioutil"
)

func main() {
	file := "/Users/liuchao56/2"
	data, _ := ioutil.ReadFile(file)

	offBytes := data[len(data)-82:]
	measurementOff := utils.BytesToInt64(offBytes[:8])
	measurementSize := utils.BytesToInt64(offBytes[8:16])
	//fmt.Println(offBytes)
	fmt.Println(measurementOff)
	fmt.Println(measurementSize)
	fmt.Println(utils.BytesToInt64(offBytes[16:24]))
	fmt.Println(utils.BytesToInt64(offBytes[24:32]))
	fmt.Println(utils.BytesToInt64(offBytes[32:40]))
	fmt.Println(utils.BytesToInt64(offBytes[40:48]))
	fmt.Println(utils.BytesToInt64(offBytes[48:56]))
	fmt.Println(utils.BytesToInt64(offBytes[56:64]))
	fmt.Println(utils.BytesToInt64(offBytes[64:72]))
	fmt.Println(utils.BytesToInt64(offBytes[72:80]))
	fmt.Println(utils.BytesToInt16(offBytes[80:82]))

	magic := data[:4]
	fmt.Println(string(magic))

	readMemBlock(data[measurementOff : measurementOff+measurementSize])

}

func readMemBlock(data []byte) {

	fmt.Println("mem block start")

	// tail
	tail := data[len(data)-66:]
	// mem Off
	memOff := utils.BytesToInt64(tail[0:8])
	memSize := utils.BytesToInt64(tail[8:16])
	fmt.Println(memOff)
	fmt.Println(memSize)
	fmt.Println(utils.BytesToInt64(tail[16:24]))
	fmt.Println(utils.BytesToInt64(tail[24:32]))
	fmt.Println(utils.BytesToInt64(tail[32:40]))
	fmt.Println(utils.BytesToInt64(tail[40:48]))
	fmt.Println(utils.BytesToInt64(tail[48:56]))
	fmt.Println(utils.BytesToInt64(tail[56:64]))
	fmt.Println(utils.BytesToInt16(tail[64:66]))

	mems := data[memOff : memOff+memSize]
	fmt.Println(data)
	fmt.Println(utils.BytesToInt64(mems[1:9]))

}
