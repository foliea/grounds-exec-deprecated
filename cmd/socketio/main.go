package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/runner"
	socketio "github.com/googollee/go-socket.io"
)

var (
	port             = flag.String("p", ":8080", "Address and port to serve")
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

	server.On("connection", func(so socketio.Socket) {
		runner := &runner.Runner{
			Client: client,
			Input:  make(chan []byte),
			Output: make(chan []byte),
			Errs:   make(chan error),
		}
		go runner.Launch()
		so.On("run", func(msg string) {
			so.Emit("run", msg)
			runner.Input <- []byte(msg)
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
		go func() {
			for out := range runner.Output {
				so.Emit("run", string(out[0:len(out)]))
			}
		}()
		go func() {
			for err := range runner.Errs {
				log.Println(err)
			}
		}()
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.HandleFunc("/socket.io/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", *authorized)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})

	log.Printf("Using docker host: %s and docker repository: %s", *dockerAddr, *dockerRepository)
	log.Printf("Authorizing: %s\n", *authorized)
	log.Printf("Listening on: %s\n", *port)
	log.Fatal(http.ListenAndServe(*port, nil))
}
