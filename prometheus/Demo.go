package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var addr1 = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	fmt.Println("add", *addr1)
	flag.Parse()
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr1, nil))
}
