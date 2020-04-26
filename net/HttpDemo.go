package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r, "r")
	fmt.Fprintf(w, "hello")
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/hello", hello)
	log.Fatal(server.ListenAndServe())
}
