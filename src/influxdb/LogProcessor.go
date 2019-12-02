package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	time int64
	data string
}

type Reader interface {
	read()
}

type Writer interface {
	write()
}
type Processor interface {
	process(rc chan string, wc chan *Message)
}
type FileReader struct {
	path     string
	readChan chan *string
}

type InfluxDBWriter struct {
	dataChan chan *Message
	server   string
}

type LogProcessor struct {
}

func (r *FileReader) read() {
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
		str := string(bytes)
		fmt.Println("reader:", str)
		r.readChan <- &str
	}
}

func (l *LogProcessor) process(rc chan *string, wc chan *Message) {
	for str := range rc {
		strArr := strings.Split(*str, "@")
		if len(strArr) != 2 {
			continue
		}
		time, _ := strconv.ParseInt(strArr[0], 10, 64)
		message := &Message{
			time: time,
			data: strArr[1],
		}
		wc <- message
	}

}

func (w *InfluxDBWriter) write() {
	for message := range w.dataChan {
		fmt.Println("writer:", *message)
	}
}

func main() {

	readChan := make(chan *string)
	writeChan := make(chan *Message)
	reader := FileReader{path: "/Users/liuchao56/log", readChan: readChan}
	process := LogProcessor{}

	writer := InfluxDBWriter{
		dataChan: writeChan,
		server:   "",
	}
	go reader.read()
	go process.process(readChan, writeChan)
	go writer.write()
	time.Sleep(500 * time.Second)
}
