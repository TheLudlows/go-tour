package main

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

func main() {
	var file = "/Users/liuchao56/.influxdb/data/testdb13/autogen/10/index/0/L0-00000001.tsl"

	//var file = "H:/influxdb-1.7.9-1/influxdb/data/mydb/autogen/3/index/4/L0-00000001.tsl"
	bytes, _ := ioutil.ReadFile(file)
	var len = 0
	fmt.Println(bytes)
	fmt.Println("flag ", bytes[:1])
	len += 1
	id, n := binary.Uvarint(bytes[len:])
	fmt.Println(id, n)
	len += n
	fmt.Println(string(bytes[len+3:]))

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
