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
	// Add a timeout call to docker version

	events := &Events{client: client, server: server}
	events.Bind()

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

type Events struct {
	client *runner.Client
	server *socketio.Server
}

func (e *Events) Bind() {
	e.server.On("connection", e.Connection)
	e.server.On("error", e.Error)
}

func (e *Events) Connection(so socketio.Socket) {
	runner := &runner.Runner{
		Client: e.client,
		Input:  make(chan []byte),
		Output: make(chan []byte),
		Errs:   make(chan error),
	}
	go runner.Launch()

	c := &Connection{runner: runner, so: so}

	so.On("run", c.Run)
	so.On("disconnection", e.Disconnection)
	go c.StreamOutput("run", runner.Output)
	go LogErrors(runner.Errs)
}

func (e *Events) Disconnection() {
	log.Println("disconnection")
}

func (e *Events) Error(so socketio.Socket, err error) {
	LogError(err)
}

type Connection struct {
	runner *runner.Runner
	so     socketio.Socket
}

func (c *Connection) Run(msg string) {
	// There is a bug with go-socket-io, this is a trick to prevent it before it gets patched upstream
	c.so.Emit("msg", msg)

	c.runner.Input <- []byte(msg)
}

func (c *Connection) StreamOutput(event string, output chan []byte) {
	for msg := range output {
		c.so.Emit(event, string(msg[0:len(msg)]))
	}
}

func LogErrors(errs chan error) {
	for err := range errs {
		LogError(err)
	}
}

func LogError(err error) {
	log.Println("error: ", err)
}
