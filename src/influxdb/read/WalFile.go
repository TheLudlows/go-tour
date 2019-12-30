package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("read")
	file, e := os.Open("C:/Users/four/.influxdb/wal/db3/autogen/7/1.wal")
	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}
	bufReader := bufio.NewReader(file)
	b, err := bufReader.ReadByte()
	if err != nil {
		log.Print(fmt.Sprint("read file", e))
	}
	fmt.Println(b)
	bs := make([]byte, 4)
	bufReader.Read(bs)
	fmt.Println(bs)

	bytebuff := bytes.NewBuffer(bs)
	var data int32
	binary.Read(bytebuff, binary.BigEndian, &data)
	//bufReader.Read(make([] byte,4))
	fmt.Println(int(data))

	data2 := make([]byte, data)
	bufReader.Read(data2)
	fmt.Println(string(data2))

	b2, _ := bufReader.ReadByte()
	fmt.Println(b2)
	bs2 := make([]byte, 4)
	bufReader.Read(bs2)
	fmt.Println(bs2)

}
