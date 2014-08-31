package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/runner"
	socketio "github.com/googollee/go-socket.io"
)

var (
	dockerAddr       = flag.String("e", "unix:///var/run/docker.sock", "Docker API endpoint")
	dockerRepository = flag.String("r", "grounds", "Docker repository to use for images")
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
		var runner = &runner.Runner{
			Client: client,
			Input:  make(chan []byte),
			Output: make(chan []byte),
			Errs:   make(chan error),
		}
		log.Println("on connect")
		go runner.Read()
		so.On("run message", func(msg string) {
			go func() {
				runner.Input <- []byte(msg)
				log.Println("msg sent")
			}()
		})
		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
		go func() {
			for out := range runner.Output {
				so.Emit("run message", out)
				log.Println("rofl")
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
		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		server.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}
