package main

import (
	"log"
	"net/http"

	"github.com/folieadrien/grounds/execcode"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func readExecAndWrite(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		execcode.Execute()
		if err = conn.WriteMessage(messageType, p); err != nil {
			return
		}
	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	readExecAndWrite(conn)
}
