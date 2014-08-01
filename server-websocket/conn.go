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

func readWriteOutputStream(conn *websocket.Conn, output io.Reader, stream string) {
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		out := Output{Stream: stream, Chunk: scanner.Text()}
		outBytes, err := json.Marshal(out); if err != nil {
			log.Println(err)
			return
		}
		if err = conn.WriteMessage(websocket.TextMessage, outBytes); err != nil {
			log.Println(err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}
}

func readExecAndWrite(conn *websocket.Conn, exec *execcode.Client) error {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		i := Input{}
		if err = json.Unmarshal(p, &i); err != nil {
			return err
		}
		if exec.IsBusy {
			if err = exec.Interrupt(); err != nil {
				return err
			}
		}
		_, err = exec.Execute(i.Language, i.Code, func (stdout, stderr io.Reader) error {
			go readWriteOutputStream(conn, stdout, "stdout")
			go readWriteOutputStream(conn, stderr, "stderr")
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
	exec, err := execcode.NewClient(*dockerAddr, *dockerRegistry) 
	if err != nil {
		log.Println(err)
		return
	}
	if err := readExecAndWrite(conn, exec); err != nil {
		log.Println(err)
		return
	}
}
