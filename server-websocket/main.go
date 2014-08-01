package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	serveAddr = flag.String("p", ":8080", "Address and port to serve")
	dockerAddr = flag.String("d", "unix:///var/run/docker.sock", "Docker host endpoint")
	dockerRegistry = flag.String("r", "", "Docker registry used for images")
)

func main() {
	flag.Parse()

	log.Printf("Listening on: %s\n", *serveAddr)
	http.HandleFunc("/ws", serveWs)
	if err := http.ListenAndServe(*serveAddr, nil); err != nil {
		log.Fatal(err)
	}
}
