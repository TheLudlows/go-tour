package main

import (
	"bufio"
	"fmt"
	"github.com/golang/snappy"
	"influxdb/utils"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("read")
	file, e := os.Open("/Users/liuchao56/.influxdb/wal/testdb13/autogen/10/1")
	if e != nil {
		log.Print(fmt.Sprint("open file", e))
	}
	bufReader := bufio.NewReader(file)

	var a = 0
	for {
		b, err := bufReader.ReadByte()
		if err == io.EOF || b == 0 {
			break
		}
		fmt.Println("flag", b)
		bs := make([]byte, 4)
		bufReader.Read(bs)
		size := utils.BytesToInt32(bs)

		data := make([]byte, size)
		bufReader.Read(data)
		a++
		var n = 0
		len, _ := snappy.DecodedLen(data)
		fmt.Println("deco size ", len)

		deComp := make([]byte, len)
		snappy.Decode(deComp, data)
		n++
		fmt.Print("type", deComp[:n])
		vType := deComp[0]
		keyLen := utils.BytesToInt16(deComp[n : n+2])
		n += 2
		fmt.Print(" key len:", keyLen)
		fmt.Print(" key:", string(deComp[n:n+int(keyLen)]))
		n += int(keyLen)
		fmt.Print(" count:", utils.BytesToInt32(deComp[n:n+4]))
		n += 4
		fmt.Print(" time:", utils.BytesToInt64(deComp[n:n+8]))
		n += 8
		switch vType {
		case 2:
			fmt.Println(" value:", utils.BytesToInt64(deComp[n:n+8]))
			break
		case 4:
			fmt.Println(" value:", string(deComp[n:n+8]))

		}

	}
	fmt.Println(a)

}
