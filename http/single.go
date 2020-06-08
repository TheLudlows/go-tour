package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"time"
)

var e1 = []byte("error=1")
var e2 = []byte("_code=")

func main() {
	var url = "http://127.0.0.1:9000/trace1.data"
	client := &http.Client{Timeout: 100 * time.Second}
	resp, _ := client.Get(url)
	start := time.Now().Unix()

	read := bufio.NewReader(resp.Body)
	Ids := make([]string, 1024*16)

	for {
		line, _, e := read.ReadLine()
		if e != nil {
			break
		}

		index := bytes.IndexByte(line, '|')
		id := line[0:index]
		if error(line) {
			Ids = append(Ids, string(id))
		}
	}
	fmt.Println(len(Ids))
	fmt.Println(time.Now().Unix() - start)
}

func error(body []byte) bool {
	from := bytes.LastIndexByte(body, '|')
	b := bytes.Index(body[from:], e1)
	if b > 0 {
		return true
	}

	c := bytes.Index(body[from:], e2)
	if b > 0 {
		return !(body[from+c+6] == '2')
	}
	return false

}
