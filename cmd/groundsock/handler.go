package main

import (
	"log"
	"net/http"

	"github.com/folieadrien/grounds/pkg/execcode"
	"github.com/gorilla/websocket"
)

type RunHandler struct {
	upgrader   *websocket.Upgrader
	execClient *execcode.Client
}

func NewRunHandler(debug bool, execClient *execcode.Client) *RunHandler {
	if execClient == nil {
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
		upgrader:   upgrader,
		execClient: execClient,
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
	c := &connection{ws: ws, execClient: h.execClient}
	c.read()
}
