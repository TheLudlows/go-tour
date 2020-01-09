package main

import (
	"fmt"
	"influxdb/utils"
	"io/ioutil"
)

func main() {
	var file = "/Users/liuchao56/.influxdb/data/testdb13/autogen/10/1"
	bytes, _ := ioutil.ReadFile(file)
	len := len(bytes)
	// footer
	indexOffset := utils.BytesToInt32(bytes[len-4:])
	fmt.Println("indexOffset ", indexOffset)

	// index
	var i = int(indexOffset)
	keyLen := utils.BytesToInt16(bytes[indexOffset : indexOffset+2])
	i += 2
	// key len
	fmt.Println(keyLen)
	// key
	fmt.Println(string(bytes[i : i+int(keyLen)]))
	i += int(keyLen)
	// value type
	fmt.Println(bytes[i : i+1])
	i += 1
	// index entry count
	fmt.Println(utils.BytesToInt16(bytes[i : i+2]))
	i += 2

	// index entry 28 byte
	entry := bytes[i : i+28]
	// min time
	fmt.Println(utils.BytesToInt64(entry[:8]))
	// max time
	fmt.Println(utils.BytesToInt64(entry[8:16]))
	// block off
	blockOffset := utils.BytesToInt64(entry[16:24])
	fmt.Println(blockOffset)
	// block size
	blockSize := utils.BytesToInt32(entry[24:28])
	fmt.Println(blockSize)
}
