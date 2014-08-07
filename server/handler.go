package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"strconv"
	"sync"

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
	dockerAddr     string
	dockerRegistry string
	debug          bool
	upgrader       *websocket.Upgrader
	conn           *websocket.Conn
	execClient     *execcode.Client
	mu             sync.Mutex
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
		dockerAddr:     dockerAddr,
		dockerRegistry: dockerRegistry,
		debug:          debug,
		upgrader:       upgrader,
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

	defer h.conn.Close()

	if h.execClient, err = execcode.NewClient(h.dockerAddr, h.dockerRegistry); err != nil {
		log.Println(err)
		return
	}
	if err := h.readExecAndSendOutput(); err != nil {
		log.Println(err)
	}
	if h.execClient.IsBusy {
		h.execClient.Interrupt()
	}
}

func (h *WsHandler) readExecAndSendOutput() error {
	for {
		input := Input{}
		if err := h.conn.ReadJSON(&input); err != nil {
			return err
		}
		// Interrupt execcode execution if already running for this client
		if h.execClient.IsBusy {
			if err := h.execClient.Interrupt(); err != nil {
				return err
			}
		}
		go func() {
			// Execute code with execcode and send output to the client
			status, err := h.execClient.Execute(input.Language, input.Code,
				func(out, err io.Reader) {
					h.sendOutput("stdout", out)
					h.sendOutput("stderr", err)
				})
			if err != nil {
				return
			}
			if !h.execClient.Interrupted() {
				h.sendResponse("status", strconv.Itoa(status))
			}
		}()
	}
}

func (h *WsHandler) sendOutput(stream string, output io.Reader) {
	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		h.sendResponse(stream, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func (h *WsHandler) sendResponse(stream, chunk string) {
	response := Response{Stream: stream, Chunk: chunk}
	h.mu.Lock()
	if err := h.conn.WriteJSON(response); err != nil {
		log.Println(err)
	}
	h.mu.Unlock()
}
