package main

import (
	"fmt"
	"go-tour/influxdb/utils"
	"io/ioutil"
)

func main() {
	file := "/Users/liuchao56/3"
	data, _ := ioutil.ReadFile(file)

	offBytes := data[len(data)-82:]
	measurementOff := utils.BytesToInt64(offBytes[:8])
	measurementSize := utils.BytesToInt64(offBytes[8:16])
	//fmt.Println(offBytes)
	fmt.Println("Measurement Block off:", measurementOff)
	fmt.Println("Measurement Block size:", measurementSize)
	fmt.Println("SeriesID Set off:", utils.BytesToInt64(offBytes[16:24]))
	fmt.Println("SeriesID Set size:", utils.BytesToInt64(offBytes[24:32]))
	fmt.Println("TombstoneSeriesIDSet off:", utils.BytesToInt64(offBytes[32:40]))
	fmt.Println("TombstoneSeriesIDSet size:", utils.BytesToInt64(offBytes[40:48]))
	fmt.Println("SeriesSketch off:", utils.BytesToInt64(offBytes[48:56]))
	fmt.Println("SeriesSketch size:", utils.BytesToInt64(offBytes[56:64]))
	fmt.Println("TombstoneSketch off:", utils.BytesToInt64(offBytes[64:72]))
	fmt.Println("TombstoneSketch size:", utils.BytesToInt64(offBytes[72:80]))
	fmt.Println("version:", utils.BytesToInt16(offBytes[80:82]))

	magic := data[:4]
	fmt.Println(string(magic))

	readMeasurementBlock(data[measurementOff : measurementOff+measurementSize])
	readTagSet(data[4:measurementOff])

}

func readMeasurementBlock(data []byte) {

	// tail
	tail := data[len(data)-66:]
	memOff := utils.BytesToInt64(tail[0:8])
	memSize := utils.BytesToInt64(tail[8:16])
	fmt.Println("Measurements off:", memOff)
	fmt.Println("Measurements size:", memSize)
	fmt.Println("Measurements HashIndex off:", utils.BytesToInt64(tail[16:24]))
	fmt.Println("Measurements HashIndex size:", utils.BytesToInt64(tail[24:32]))
	fmt.Println("Measurements Sketch off:", utils.BytesToInt64(tail[32:40]))
	fmt.Println("Measurements Sketch size:", utils.BytesToInt64(tail[40:48]))
	fmt.Println("Measurements TombStone Sketch off:", utils.BytesToInt64(tail[48:56]))
	fmt.Println("Measurements TombStone Sketch size:", utils.BytesToInt64(tail[56:64]))
	fmt.Println("version:", utils.BytesToInt16(tail[64:66]))

	mems := data[memOff : memOff+memSize]
	// read a measurement
	fmt.Println(mems[0:1])

}
func readTagSet(data []byte) {

	tail := data[len(data)-58:]
	fmt.Println(tail)
	fmt.Println("value off:", utils.BytesToInt64(tail[0:8]))
	fmt.Println("value size:", utils.BytesToInt64(tail[8:16]))
	fmt.Println("key off:", utils.BytesToInt64(tail[16:24]))
	fmt.Println("key size:", utils.BytesToInt64(tail[24:32]))
	fmt.Println("hash off:", utils.BytesToInt64(tail[32:40]))
	fmt.Println("hash size:", utils.BytesToInt64(tail[40:48]))
	fmt.Println("total size:", utils.BytesToInt64(tail[48:56]))
	fmt.Println("version:", tail[56:58])

}
