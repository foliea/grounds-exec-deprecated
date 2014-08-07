package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	serveAddr      = flag.String("p", ":8080", "Address and port to serve")
	dockerAddr     = flag.String("e", "unix:///var/run/docker.sock", "Docker API endpoint")
	dockerRegistry = flag.String("r", "grounds", "Docker registry used for images")
	debug          = flag.Bool("d", false, "Debug mode")
)

func main() {
	flag.Parse()

	if *debug {
		log.Printf("Warning: using debug mode, origin check disabled")
	}
	log.Printf("Using docker host: %s and docker registry: %s", *dockerAddr, *dockerRegistry)
	log.Printf("Listening on: %s\n", *serveAddr)

	http.Handle("/ws", NewWsHandler(*debug, *dockerAddr, *dockerRegistry))
	if err := http.ListenAndServe(*serveAddr, nil); err != nil {
		log.Fatal(err)
	}
}
