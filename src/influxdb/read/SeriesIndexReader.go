package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	file, e := os.Open("/Users/liuchao56/.influxdb/data/testdb12/_series/07/index")
	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}
	bufReader := bufio.NewReader(file)

	magic := make([]byte, 4)
	bufReader.Read(magic)
	fmt.Println(string(magic))
	data := make([]byte, 65)
	bufReader.Read(data)
	fmt.Println(data[:1])
	fmt.Println(bytesToInt64(data[1:9]))
	fmt.Println(bytesToInt64(data[9:17]))
	fmt.Println(bytesToInt64(data[17:25]))
	fmt.Println(bytesToInt64(data[25:33]))
	fmt.Println(bytesToInt64(data[33:41]))
	fmt.Println(bytesToInt64(data[41:49]))
	fmt.Println(bytesToInt64(data[49:57]))
	fmt.Println(bytesToInt64(data[57:65]))
}
func bytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}
