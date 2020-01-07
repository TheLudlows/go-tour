package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"log"
	"os"
)

func main() {
	//file, e := os.Open("/Users/liuchao56/.influxdb/data/testdb12/autogen/4/index/0/L0-00000001.tsl")
	file, e := os.Open("/Users/liuchao56/L0-00000002.tsl")

	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}

	bufReader := bufio.NewReader(file)
	line, _, _ := bufReader.ReadLine()

	fmt.Println(len(line))

	//return
	fmt.Println(line)
	length := 0
	flag := line[0:1]
	fmt.Println(flag)
	length++

	id, n := binary.Uvarint(line[length:])
	fmt.Println(id, n)
	length += n
	fmt.Println(string(line[3:9]))

	/*len := BytesToInt32(mBytes)
	lenBytes := make([] byte,len)
	bufReader.Read(lenBytes)
	fmt.Println(string(lenBytes))

	flag1,_ := bufReader.ReadByte()
	fmt.Println(flag1)

	idBytes1 := make([]byte,4)
	bufReader.Read(idBytes1)
	fmt.Println(biu.BytesToBinaryString(idBytes1))*/

}
