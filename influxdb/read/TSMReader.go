package main

import (
	"encoding/binary"
	"fmt"
	"github.com/jwilder/encoding/simple8b"
	"go-tour/influxdb/utils"
	"io/ioutil"
	"math"
)

func main() {
	var file = "/Users/liuchao56/.influxdb/data/testdb13/autogen/10/1"
	bytes, _ := ioutil.ReadFile(file)
	len := len(bytes)
	// footer
	indexOffset := utils.BytesToInt32(bytes[len-4:])
	fmt.Println("indexOffset:", indexOffset)

	// index
	var i = int(indexOffset)
	keyLen := utils.BytesToInt16(bytes[indexOffset : indexOffset+2])
	i += 2
	// key len
	fmt.Println("keyLen:", keyLen)
	// key
	fmt.Println("series key:", string(bytes[i:i+int(keyLen)]))
	i += int(keyLen)
	// value type
	fmt.Println("value type:", bytes[i:i+1])
	i += 1
	// index entry count
	fmt.Println("index entry count:", utils.BytesToInt16(bytes[i:i+2]))
	i += 2

	// index entry 28 byte
	entry := bytes[i : i+28]
	// min time
	fmt.Println("min time:", utils.BytesToInt64(entry[:8]))
	// max time
	fmt.Println("max time:", utils.BytesToInt64(entry[8:16]))
	// block off
	blockOffset := utils.BytesToInt64(entry[16:24])
	fmt.Println("block off:", blockOffset)
	// block size
	blockSize := utils.BytesToInt32(entry[24:28])
	fmt.Println("block size:", blockSize)

	//read block
	blockData := bytes[blockOffset : blockOffset+int64(blockSize)]
	fmt.Println(blockData)

	fmt.Println("crc num :", utils.BytesToInt32(blockData[0:4]))

	fmt.Println("value type:", blockData[4])
	blockData = blockData[5:]
	timeLen, n := binary.Uvarint(blockData)
	fmt.Println("time Len:", timeLen, "time Len size:", n)
	timeData := blockData[int64(n) : int64(n)+int64(timeLen)]
	fmt.Println(timeData)
	valueData := blockData[int64(n)+int64(timeLen):]
	fmt.Println(valueData)
	fmt.Println("encode type:", timeData[0])

	log10 := int(timeData[0] & 0xf)
	fmt.Println(math.Pow10(log10))

	fmt.Println("time stamp:", utils.BytesToInt64(timeData[1:9]))
	timeData = timeData[9:]
	decoder := simple8b.NewDecoder(timeData)

	out := make([]uint64, 0)
	for decoder.Next() {
		out = append(out, decoder.Read())
	}
	fmt.Println(out)

}
