package main

import (
	"bufio"
	"fmt"
	"go-tour/influxdb/utils"
	"io"
	"log"
	"os"
)

func main() {
	//file, e := os.Open("/Users/liuchao56/.influxdb/data/testdb13/_series/00/0000")
	file, e := os.Open("H:/influxdb-1.7.9-1/influxdb/data/db1/_series/04/0000")
	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}
	bufReader := bufio.NewReader(file)

	magic := make([]byte, 4)
	bufReader.Read(magic)
	fmt.Println("magic:", string(magic))

	b, _ := bufReader.ReadByte()
	fmt.Println(b)
	for {
		flag, err := bufReader.ReadByte()
		if err == io.EOF {
			fmt.Println("read over !")
			return
		}
		if flag == 0 {
			fmt.Println("read over !", flag)
			return
		}
		fmt.Println("flag:", flag)

		idBytes := make([]byte, 8)
		bufReader.Read(idBytes)
		fmt.Println("series id", utils.BytesToInt64(idBytes))

		sizeByte, _ := bufReader.ReadByte()
		fmt.Println("series key size:", sizeByte)

		key := make([]byte, sizeByte)
		bufReader.Read(key)
		fmt.Println("series key:", string(key))

		/*measurementSizeBytes := key[:2]
		measurementSize := utils.BytesToInt16(measurementSizeBytes)
		fmt.Println(measurementSize)
		measurement := key[2 : 2+measurementSize]
		fmt.Println(string(measurement))
		tagSize := key[9:10]
		fmt.Println(tagSize)
		tagKeySize := key[10:12]
		fmt.Println(utils.BytesToInt16(tagKeySize))
		tagKey := key[12 : 12+5]
		fmt.Println(string(tagKey))
		tagValueSizeByte := key[17:19]
		tagValueSize := utils.BytesToInt16(tagValueSizeByte)
		fmt.Println(tagValueSize)
		tagValue := key[19 : 19+tagValueSize]
		fmt.Println(string(tagValue))*/
	}

}
