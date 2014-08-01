package main

import (
	"log"
	"encoding/json"
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

type Input struct {
    Language string `json:"language"`
    Code string `json:"code"`
}

type Output struct {
    Stream string `json:"stream"`
    Chunk string `json:"chunk"`
}

func readExecAndWrite(conn *websocket.Conn, exec *execcode.Client) error {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		i := Input{}
		if err = json.Unmarshal(p, &i); err != nil {
			return err
		}
		_, err = exec.Execute(i.Language, i.Code, func (stdout, stderr io.Reader) error {
			// Fix: Close readers
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				out := Output{Stream: "stdout", Chunk: scanner.Text()}
				outBytes, err := json.Marshal(out); if err != nil {
					return err
				}
				if err = conn.WriteMessage(messageType, outBytes); err != nil {
					return err
				}
			}
			if err := scanner.Err(); err != nil {
				return err
			}
			log.Println("looping twice")
			return nil
		})
		log.Println("looping twice")
		if err != nil {
			return err
		}
		log.Println("looping twice")
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
	exec, err := execcode.NewClient(*dockerAddr, *dockerRegistry) 
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("New client, using docker host: %s and docker registry: %s", *dockerAddr, *dockerRegistry)
	if err := readExecAndWrite(conn, exec); err != nil {
		log.Println(err)
		return
	}
}
