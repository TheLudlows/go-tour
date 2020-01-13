package main

import (
	"bufio"
	"fmt"
	"go-tour/influxdb/utils"
	"log"
	"os"
)

func main() {
	file, e := os.Open("/Users/liuchao56/.influxdb/data/testdb12/_series/03/index")
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
	fmt.Println(utils.BytesToInt64(data[1:9]))
	fmt.Println(utils.BytesToInt64(data[9:17]))
	fmt.Println(utils.BytesToInt64(data[17:25]))
	fmt.Println(utils.BytesToInt64(data[25:33]))
	fmt.Println(utils.BytesToInt64(data[33:41]))
	fmt.Println(utils.BytesToInt64(data[41:49]))
	fmt.Println(utils.BytesToInt64(data[49:57]))
	fmt.Println(utils.BytesToInt64(data[57:65]))

	//keyOffset -> id
	keyIDMapBlockSize := utils.BytesToInt64(data[41:49])
	fmt.Println("keyIDMapBlockSize", keyIDMapBlockSize)
	keyIDMapData := make([]byte, keyIDMapBlockSize)
	fmt.Println("keyIDMapData len", len(keyIDMapData))
	bufReader.Read(keyIDMapData)
	fmt.Println(utils.BytesToInt32(keyIDMapData[:4]))
	fmt.Println(utils.BytesToInt32(keyIDMapData[4:8]))
	fmt.Println(utils.BytesToInt64(keyIDMapData[8:16]))

	fmt.Println(utils.BytesToInt32(keyIDMapData[16:20]))
	fmt.Println(utils.BytesToInt32(keyIDMapData[20:24]))
	fmt.Println(utils.BytesToInt64(keyIDMapData[24:32]))

	IDOffsetMapBlockSize := utils.BytesToInt64(data[57:65])
	IDOffsetMapData := make([]byte, IDOffsetMapBlockSize)
	n, err := bufReader.Read(IDOffsetMapData)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println("read size", n)
	fmt.Println(IDOffsetMapData[:8])
	fmt.Println(utils.BytesToInt64(IDOffsetMapData[:8]))
	fmt.Println(utils.BytesToInt32(IDOffsetMapData[8:12]))
	fmt.Println(utils.BytesToInt32(IDOffsetMapData[12:16]))

	fmt.Println(utils.BytesToInt64(IDOffsetMapData[16:24]))
	fmt.Println(utils.BytesToInt32(IDOffsetMapData[24:28]))
	fmt.Println(utils.BytesToInt32(IDOffsetMapData[28:32]))

}
