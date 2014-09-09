package handler

import (
	"log"
	"time"

	"github.com/folieadrien/grounds/pkg/runner"
	socketio "github.com/googollee/go-socket.io"
)

type Handler struct {
	Client *runner.Client
	Server *socketio.Server
}

func (h *Handler) Bind() {
	h.Server.On("connection", h.NewConnection)
	h.Server.On("error", h.Error)
}

func (h *Handler) NewConnection(so socketio.Socket) {
	log.Println("new connection")

	runner := &runner.Runner{
		Client: h.Client,
		Input:  make(chan []byte),
		Output: make(chan []byte),
		Errs:   make(chan error),
	}
	go runner.Watch()

	c := &Connection{
		input:  runner.Input,
		output: runner.Output,
		time: time.Now(),
		event:  "run",
		so:     so,
	}
	go c.Write()

	so.On("run", c.Read)
	so.On("disconnection", h.Disconnection)

	go LogErrors(runner.Errs)
}

func (h *Handler) Disconnection() {
	log.Println("disconnection")
}

func (h *Handler) Error(so socketio.Socket, err error) {
	LogError(err)
}

func LogErrors(errs chan error) {
	for err := range errs {
		LogError(err)
	}
}

func LogError(err error) {
	log.Println("error: ", err)
}
