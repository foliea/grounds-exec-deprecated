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

func readExecAndWrite(conn *websocket.Conn, exec *execcode.Client) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		err = exec.Execute("ruby", "puts \"lol\"", func() error {
			if err := conn.WriteMessage(messageType, p); err != nil {
					return err
			}
			return nil
		})
		if err != nil {
			return err
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
	exec, err := execcode.NewClient("http://178.62.34.175:4243") 
	if err != nil {
		log.Println(err)
		return
	}
	if err := readExecAndWrite(conn, exec); err != nil {
		log.Println(err)
		return
	}
}
