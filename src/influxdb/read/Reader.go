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
	file, e := os.Open("C:/Users/four/.influxdb/data/db3/_series/05/0000")
	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}
	bufReader := bufio.NewReader(file)

	/*line,_,_ := bufReader.ReadLine();
	fmt.Println(string(line))*/

	magic := make([]byte, 4)
	bufReader.Read(magic)
	fmt.Println(string(magic))

	b, _ := bufReader.ReadByte()
	fmt.Println(b)

	flag, _ := bufReader.ReadByte()
	fmt.Println(flag)

	idBytes := make([]byte, 8)
	bufReader.Read(idBytes)
	bytebuff := bytes.NewBuffer(idBytes)
	var seriesId int64
	binary.Read(bytebuff, binary.BigEndian, &seriesId)
	fmt.Println(seriesId)

	sizeByte := make([]byte, 1)
	bufReader.Read(sizeByte)
	sizeBuf := bytes.NewBuffer(sizeByte)
	var size uint8
	binary.Read(sizeBuf, binary.BigEndian, &size)
	fmt.Println(size)

	body := make([]byte, size)
	bufReader.Read(body)
	fmt.Println(string(body))

	b1, _ := bufReader.ReadByte()
	fmt.Println(b1)

	idBytes2 := make([]byte, 8)
	bufReader.Read(idBytes2)
	bytebuff2 := bytes.NewBuffer(idBytes2)
	var seriesId2 int64
	binary.Read(bytebuff2, binary.BigEndian, &seriesId2)
	fmt.Println(seriesId2)

	sizeByte1 := make([]byte, 1)
	bufReader.Read(sizeByte1)
	sizeBuf1 := bytes.NewBuffer(sizeByte1)
	var size1 uint8
	binary.Read(sizeBuf1, binary.BigEndian, &size1)
	fmt.Println(size1)

	bufReader.Read(body)
	fmt.Println(string(body))

}
