package handler

import (
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/runner"
	"github.com/gorilla/websocket"
)

type RunHandler struct {
	upgrader *websocket.Upgrader
	client   *runner.Client
}

func NewRunHandler(debug bool, dockerAddr, dockerRegistry string) *RunHandler {
	client, err := runner.NewClient(dockerAddr, dockerRegistry)
	if err != nil {
		return nil
	}
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if debug {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}
	return &RunHandler{
		upgrader: upgrader,
		client:   client,
	}
}

func (h *RunHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	ws, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	runner := &runner.Runner{
		Client: h.client,
		Input:  make(chan []byte),
		Output: make(chan []byte),
		Errs:   make(chan error),
	}
	conn := &connection{
		ws:      ws,
		receive: runner.Input,
		send:    runner.Output,
	}
	go runner.Read()
	go conn.reader()
	go conn.writer()

	for err := range runner.Errs {
		log.Println(err)
	}
}
