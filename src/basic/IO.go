package main

import (
	"fmt"
	"io"
	"os"
)

func readFormReader(reader io.Reader) []byte {
	bytes := make([]byte, 1000)
	n, _ := reader.Read(bytes)
	return bytes[:n]
}
func main() {
	//bytes := readFormReader(strings.NewReader("four hello"))
	//bytes := readFormReader(os.Stdin)

	file, error := os.Open("IO.go")
	//defer file.Close()
	bytes := readFormReader(file)
	fmt.Println(error)
	fmt.Println(bytes)

}
