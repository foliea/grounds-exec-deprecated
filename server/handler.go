package main

import (
	"log"
	"net/http"

	"github.com/folieadrien/grounds/execcode"
	"github.com/gorilla/websocket"
)

type WsHandler struct {
	debug      bool
	upgrader   *websocket.Upgrader
	execClient *execcode.Client
}

func NewWsHandler(debug bool, execClient *execcode.Client) *WsHandler {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	if debug {
		upgrader.CheckOrigin = func(r *http.Request) bool {
			return true
		}
	}
	return &WsHandler{
		debug:      debug,
		upgrader:   upgrader,
		execClient: execClient,
	}
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
