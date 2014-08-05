package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/folieadrien/grounds/execcode"
	"github.com/gorilla/websocket"
)

type Input struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Response struct {
	Stream string `json:"stream"`
	Chunk  string `json:"chunk"`
}

type WsHandler struct {
	upgrader       *websocket.Upgrader
	conn           *websocket.Conn
	execClient     *execcode.Client
	dockerAddr     string
	dockerRegistry string
	debug          bool
}

func NewWsHandler(debug bool, dockerAddr, dockerRegistry string) *WsHandler {
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
		upgrader:       upgrader,
		dockerAddr:     dockerAddr,
		dockerRegistry: dockerRegistry,
		debug:          debug,
	}
}

func (h *WsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	var err error
	if h.conn, err = h.upgrader.Upgrade(w, r, nil); err != nil {
		log.Println(err)
		return
	}
	if h.execClient, err = execcode.NewClient(h.dockerAddr, h.dockerRegistry); err != nil {
		log.Println(err)
		return
	}
	if err := h.readExecAndSendOutput(); err != nil {
		log.Println(err)
		return
	}
}

func (h *WsHandler) readExecAndSendOutput() error {
	for {
		_, message, err := h.conn.ReadMessage()
		if err != nil {
			return err
		}
		input := Input{}
		if err = json.Unmarshal(message, &input); err != nil {
			return err
		}
		// Interrupt execcode execution if already running by the client
		if h.execClient.IsBusy {
			if err = h.execClient.Interrupt(); err != nil {
				return err
			}
		}
		// Execute code with execcode and send output to the client
		status, err := h.execClient.Execute(input.Language, input.Code,
			func(out, err io.Reader) error {
				go h.sendOutputStream("stdout", out)
				go h.sendOutputStream("stderr", err)
				return nil
			})
		if err != nil {
			return err
		}
		// Send status returned from executed program
		if err = h.sendResponse("status", strconv.Itoa(status)); err != nil {
			return err
		}
	}
}

func (h *WsHandler) sendOutputStream(stream string, output io.Reader) {
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		if err := h.sendResponse(stream, scanner.Text()); err != nil {
			log.Println(err)
			return
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func (h *WsHandler) sendResponse(stream, chunk string) error {
	response := Response{Stream: stream, Chunk: chunk}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return err
	}
	if err = h.conn.WriteMessage(websocket.TextMessage, jsonResponse); err != nil {
		return err
	}
	return nil
}
