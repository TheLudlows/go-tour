package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var e1 = []byte("error=1")
var e2 = []byte("_code=")
var client = http.Client{Timeout: time.Second * 180}

var threadGroup = sync.WaitGroup{}

var threads int

func init() {
	//每个线程下载文件的大小
	threads = 4
}
func main() {
	var url = "http://127.0.0.1:9000/trace1.data"
	fmt.Println(time.Now().Unix())
	Download(url)
	fmt.Println(time.Now().Unix())
}

func Download(url string) {

	//HEAD 方法请求服务端是否支持多线程下载,并获取文件大小
	if request, e := http.NewRequest("HEAD", url, nil); e == nil {
		if response, i := client.Do(request); i == nil {
			defer response.Body.Close()
			//得到文件大小
			ContentLength := response.ContentLength
			dispSliceDownload(ContentLength, url)
		}
	}
}

func dispSliceDownload(ContentLength int64, url string) int {

	bucket := ContentLength / int64(threads)
	//分配下载线程
	for i := 0; i < threads; i++ {
		//计算每个线程下载的区间,起始位置
		var start int64
		var end int64
		start = int64(i) * bucket
		end = int64(i+1)*bucket - 1
		if i == threads-1 {
			end = ContentLength - 1
		}
		//构建请求
		if req, e := http.NewRequest("GET", url, nil); e == nil {
			req.Header.Set(
				"Range",
				"bytes="+strconv.FormatInt(start, 10)+"-"+strconv.FormatInt(end, 10))
			threadGroup.Add(1)
			go sliceDownload(req)
		} else {
			panic(e)
		}

	}
	//等待所有线程完成下载
	threadGroup.Wait()
	return 0
}

func sliceDownload(request *http.Request) {
	defer threadGroup.Done()
	if response, e := client.Do(request); e == nil && response.StatusCode == 206 {
		defer response.Body.Close()
		read(response)
	} else {
		panic(e)
	}
}

func read(r *http.Response) {

	read := bufio.NewReader(r.Body)
	Ids := make([]string, 1024*16)

	for {
		line, _, e := read.ReadLine()
		if e != nil {
			break
		}

		index := bytes.IndexByte(line, '|')
		if index < 0 {
			continue
		}
		id := line[0:index]

		if error1(line) {
			Ids = append(Ids, string(id))
		}
	}
	fmt.Println(len(Ids))
}

func error1(body []byte) bool {
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
