package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/handler"
	"github.com/folieadrien/grounds/pkg/runner"
	socketio "github.com/googollee/go-socket.io"
)

var (
	serveAddr        = flag.String("p", ":8080", "Address and port to serve")
	dockerAddr       = flag.String("e", "unix:///var/run/docker.sock", "Docker API endpoint")
	dockerRepository = flag.String("r", "grounds", "Docker repository to use for images")
	authorized       = flag.String("a", "http://127.0.0.1:3000", "Authorized client")
)

func main() {
	flag.Parse()

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	client, err := runner.NewClient(*dockerAddr, *dockerRepository)
	if err != nil {
		log.Fatal(err)
	}

	handler := &handler.Handler{Client: client, Server: server}
	handler.Bind()

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", *authorized)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})

	log.Printf("Using docker host: %s and docker repository: %s", *dockerAddr, *dockerRepository)
	log.Printf("Authorizing: %s\n", *authorized)
	log.Printf("Listening on: %s\n", *serveAddr)
	log.Fatal(http.ListenAndServe(*serveAddr, nil))
}
