package main

import (
	"bufio"
	"fmt"
	"go-tour/influxdb/utils"
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
	fmt.Println(utils.BytesToInt64(offBytes[:8]))
	fmt.Println(utils.BytesToInt64(offBytes[8:16]))
	fmt.Println(utils.BytesToInt64(offBytes[16:24]))
	fmt.Println(utils.BytesToInt64(offBytes[24:32]))
	fmt.Println(utils.BytesToInt64(offBytes[32:40]))
	fmt.Println(utils.BytesToInt64(offBytes[40:48]))
	fmt.Println(utils.BytesToInt64(offBytes[48:56]))
	fmt.Println(utils.BytesToInt64(offBytes[56:64]))
	fmt.Println(utils.BytesToInt64(offBytes[64:72]))
	fmt.Println(utils.BytesToInt64(offBytes[72:80]))
	fmt.Println(utils.BytesToInt64(offBytes[80:82]))

}
