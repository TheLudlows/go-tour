package main

import (
	"bufio"
	"fmt"
	"influxdb/influx"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Message struct {
	t     int64
	data  string
	value string
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
	client   *influx.InfluxDB
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
		log.Println("reader:", str)
		r.readChan <- &str
	}
}

func (l *LogProcessor) process(rc chan *string, wc chan *Message) {
	for str := range rc {
		strArr := strings.Split(*str, "@")
		if len(strArr) != 3 {
			continue
		}
		time, _ := strconv.ParseInt(strArr[0], 10, 64)
		message := &Message{
			t:     time,
			data:  strArr[1],
			value: strArr[2],
		}
		wc <- message
	}

}

func (w *InfluxDBWriter) write() {
	tags := map[string]string{}
	field := map[string]interface{}{}
	for message := range w.dataChan {
		tags["value"] = message.value
		field["data"] = message.data
		field["t"] = message.t
		w.client.Insert(&tags, &field)
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
		client:   initDBClient(),
	}
	go mock_data()
	go reader.read()
	go process.process(readChan, writeChan)
	go writer.write()
	time.Sleep(500 * time.Second)
}

func initDBClient() *influx.InfluxDB {
	dbClient := influx.InfluxDB{
		Url:       "http://127.0.0.1:8086",
		Name:      "admin",
		Pwd:       "",
		Db:        "logprocess",
		Mmt:       "log",
		Precision: "ms",
	}
	dbClient.Build()
	return &dbClient
}

func mock_data() {

	file, err := os.OpenFile("/Users/liuchao56/log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(fmt.Sprintf("open file error:%s", err))
	}
	file.Seek(0, 2)
	rand.Seed(6)
	for {
		n, err := file.WriteString(fmt.Sprint(time.Now().Unix(), "@", rand.Int(), "@", rand.Int()))
		if err != nil {
			log.Print("write data error ", err)
			return
		}
		log.Print("write data total ", n)
		time.Sleep(500 * time.Millisecond)
	}
}
