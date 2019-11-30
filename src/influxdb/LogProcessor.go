package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

type Reader interface {
	read()
}

type Writer interface {
	write()
}
type Processor interface {
	process()
}
type FileReader struct {
	path string
}

type InfluxDBWriter struct {
	dataChan chan []byte
	server   string
}

type LogProcessor struct {
	dataChan chan []byte
}

func (r *FileReader) read(readChan chan []byte) {
	file, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err))
	}
	file.Seek(0, 2)
	for {
		bufReader := bufio.NewReader(file)
		bytes, _, err := bufReader.ReadLine()
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("readFile error:%s", err))
		}
		fmt.Println(string(bytes))
		readChan <- bytes
	}
}

func (w *InfluxDBWriter) write() {

}

func (l *LogProcessor) process() {

}

func main() {

	readChan := make(chan []byte)
	reader := FileReader{path: "/Users/liuchao56/log"}
	go reader.read(readChan)
	<-readChan

	time.Sleep(500 * time.Second)
}
