package main

import (
	"flag"
	"log"
	"net/http"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")

func main() {
	flag.Parse()
	log.Println("Starting web server at http://0.0.0.0:" + *port)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*root))))
}
