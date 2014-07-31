package main

import (
	"log"
	"net/http"
	"io"
	"bufio"

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
		messageType, _, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		err = exec.Execute("ruby", "3.times do\\nputs \"lol\"\\nsleep 3\\nend", func (stdout, stderr io.Reader) error {
			// Fix: Close readers
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				if err = conn.WriteMessage(messageType, scanner.Bytes()); err != nil {
					return err
				}
			}
			if err := scanner.Err(); err != nil {
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
	exec, err := execcode.NewClient("http://178.62.34.175:4243", "foliea") 
	if err != nil {
		log.Println(err)
		return
	}
	if err := readExecAndWrite(conn, exec); err != nil {
		log.Println(err)
		return
	}
}
