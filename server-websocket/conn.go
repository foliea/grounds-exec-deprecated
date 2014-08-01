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

type Input struct {
    Language string `json:"language"`
    Code string `json:"code"`
}

type Output struct {
    Stream string `json:"stream"`
    Chunk string `json:"chunk"`
}

type WsHandler struct {
	upgrader *websocket.Upgrader
	conn	*websocket.Conn
	execClient	*execcode.Client
}

func NewWsHandler() *WsHandler {
	upgrader := &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
	}
	return &WsHandler{upgrader: upgrader}
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	var err error
	h.conn, err = h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	h.execClient, err = execcode.NewClient(*dockerAddr, *dockerRegistry) 
	if err != nil {
		log.Println(err)
		return
	}
	if err := h.readExecAndWrite(); err != nil {
		log.Println(err)
		return
	}
}

func (h *WsHandler) readExecAndWrite() error {
	for {
		_, p, err := h.conn.ReadMessage()
		if err != nil {
			return err
		}
		i := Input{}
		if err = json.Unmarshal(p, &i); err != nil {
			return err
		}
		if h.execClient.IsBusy {
			if err = h.execClient.Interrupt(); err != nil {
				return err
			}
		}
		_, err = h.execClient.Execute(i.Language, i.Code, func (stdout, stderr io.Reader) error {
			go h.sendOutputStream(stdout, "stdout")
			go h.sendOutputStream(stderr, "stderr")
			return nil
		})
		if err != nil {
			return err
		}
	}
}

func (h *WsHandler) sendOutputStream(output io.Reader, stream string) {
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		output := Output{Stream: stream, Chunk: scanner.Text()}
		response, err := json.Marshal(output); if err != nil {
			log.Println(err)
			return
		}
		if err = h.conn.WriteMessage(websocket.TextMessage, response); err != nil {
			log.Println(err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}
}
