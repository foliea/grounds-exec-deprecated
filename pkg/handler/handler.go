package handler

import (
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/runner"
	"github.com/gorilla/websocket"
)

type RunHandler struct {
	upgrader *websocket.Upgrader
	runner   *runner.Client
}

func NewRunHandler(debug bool, dockerAddr, dockerRegistry string) *RunHandler {
	runner, err := runner.NewClient(dockerAddr, dockerRegistry)
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
		runner:   runner,
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
	c := &connection{ws: ws, runner: h.runner}
	c.read()
}
