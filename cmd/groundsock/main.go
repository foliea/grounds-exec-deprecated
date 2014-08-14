package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/handler"
)

var (
	serveAddr      = flag.String("p", ":8080", "Address and port to serve")
	dockerAddr     = flag.String("e", "unix:///var/run/docker.sock", "Docker API endpoint")
	dockerRegistry = flag.String("r", "grounds", "Docker registry to use for images")
	debug          = flag.Bool("d", false, "Debug mode")
)

func main() {
	flag.Parse()

	if *debug {
		log.Printf("Warning: using debug mode, origin check disabled")
	}
	log.Printf("Using docker host: %s and docker registry: %s", *dockerAddr, *dockerRegistry)

	run := handler.NewRunHandler(*debug, *dockerAddr, *dockerRegistry)
	if run == nil {
		log.Fatalf("Impossible to create run handler, verify your docker endpoint.")
	}
	http.Handle("/run", run)

	log.Printf("Listening on: %s\n", *serveAddr)
	if err := http.ListenAndServe(*serveAddr, nil); err != nil {
		log.Fatal(err)
	}
}
