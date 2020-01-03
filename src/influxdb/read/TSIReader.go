package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	file, e := os.Open("/Users/liuchao56/L1-00000001.tsi")

	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}

	file.Seek(-82, 2)
	bufReader := bufio.NewReader(file)

	offBytes := make([]byte, 82)
	bufReader.Read(offBytes)

	fmt.Println(offBytes)
	fmt.Println(BytesToInt64(offBytes[:8]))
	fmt.Println(BytesToInt64(offBytes[8:16]))
	fmt.Println(BytesToInt64(offBytes[16:24]))
	fmt.Println(BytesToInt64(offBytes[24:32]))
	fmt.Println(BytesToInt64(offBytes[32:40]))
	fmt.Println(BytesToInt64(offBytes[40:48]))
	fmt.Println(BytesToInt64(offBytes[48:56]))
	fmt.Println(BytesToInt64(offBytes[56:64]))
	fmt.Println(BytesToInt64(offBytes[64:72]))
	fmt.Println(BytesToInt64(offBytes[72:80]))
	fmt.Println(bytesToInt16(offBytes[80:82]))

}
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func bytesToInt16(buf []byte) int16 {
	return int16(binary.BigEndian.Uint16(buf))
}
