package main

import (
	"encoding/binary"
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
	//readTagSet(data[4:measurementOff])

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

	memData := data[memOff : memOff+memSize]
	fmt.Println("Measurements data len:", len(memData))
	//read hashIndex
	indexOff := utils.BytesToInt64(tail[16:24])
	indexSize := utils.BytesToInt64(tail[24:32])
	indexData := data[indexOff : indexOff+indexSize]

	fmt.Println("index data len:", len(indexData))
	fmt.Println("Measurements count:", utils.BytesToInt64(indexData[:8]))
	fmt.Println("Measurement 1 off:", utils.BytesToInt64(indexData[8:16]))
	fmt.Println("Measurement 2 off:", utils.BytesToInt64(indexData[16:24]))

	//read Measurement
	mem2Off := utils.BytesToInt64(indexData[16:24])
	mem2Data := memData[mem2Off:]
	var flag byte
	flag, mem2Data = mem2Data[0], mem2Data[1:]
	fmt.Println("Measurement 2 flag:", flag)

	fmt.Println("tag off:", utils.BytesToInt64(mem2Data[:8]))
	fmt.Println("tag size:", utils.BytesToInt64(mem2Data[8:16]))
	mem2Data = mem2Data[16:]
	//name
	size, n := binary.Uvarint(mem2Data)
	mem2Data = mem2Data[n:]
	fmt.Println("Measurement 2 name size:", size)
	fmt.Println("Measurement 2 name:", string(mem2Data[:size]))
	mem2Data = mem2Data[size:]

	// series count
	count, n := binary.Uvarint(mem2Data)
	fmt.Println("series  count:", count)
	mem2Data = mem2Data[n:]
	// series size
	seriesSize, n := binary.Uvarint(mem2Data)
	fmt.Println("series Size:", seriesSize)
	mem2Data = mem2Data[n:]

}
func readTagSet(data []byte) {

	tail := data[len(data)-58:]
	fmt.Println("value off:", utils.BytesToInt64(tail[0:8]))
	fmt.Println("value size:", utils.BytesToInt64(tail[8:16]))
	fmt.Println("key off:", utils.BytesToInt64(tail[16:24]))
	fmt.Println("key size:", utils.BytesToInt64(tail[24:32]))
	fmt.Println("hash off:", utils.BytesToInt64(tail[32:40]))
	fmt.Println("hash size:", utils.BytesToInt64(tail[40:48]))
	fmt.Println("total size:", utils.BytesToInt64(tail[48:56]))
	fmt.Println("version:", tail[56:58])

	data = readTagValue(data[1:])
	start := len(data)

	fmt.Println("read ", len(data)-start)
}
func readTagValue(data []byte) []byte {
	var flag byte
	flag, data = data[0], data[1:]
	fmt.Println("flag:", flag)

	var n int
	var size uint64
	size, n = binary.Uvarint(data)
	fmt.Println("size:", size)
	data = data[n:]
	fmt.Println("tag value:", string(data[:size]))
	data = data[size:]

	var count uint64
	count, n = binary.Uvarint(data)
	data = data[n:]
	fmt.Println("series count", count)

	var seriesSize uint64
	seriesSize, n = binary.Uvarint(data)
	data = data[n:]
	fmt.Println("series size", seriesSize)
	fmt.Println(string(data[:seriesSize]))
	data = data[seriesSize:]
	return data
}
