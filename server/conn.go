package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"sync"

	"github.com/folieadrien/grounds/execcode"
	"github.com/gorilla/websocket"
)

type connection struct {
	sync.Mutex
	ws         *websocket.Conn
	execClient *execcode.Client
}

type Input struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

type Response struct {
	Stream string `json:"stream"`
	Chunk  string `json:"chunk"`
}

// FIXME: IF CODE > 1500 program too large
// Prepare a new container
func (c *connection) read() {
	defer c.ws.Close()
	var (
		containerID string
		interrupted chan bool
	)
	for {
		var (
			input = Input{}
			err   = c.ws.ReadJSON(&input)
		)
		if containerID != "" {
			go c.interrupt(containerID, interrupted)
		}
		if err != nil {
			log.Println(err)
			return
		}
		if containerID, err = c.execClient.Prepare(input.Language, input.Code); err != nil {
			log.Println(err)
			continue
		}
		interrupted = make(chan bool, 3)
		go c.exec(containerID, interrupted)
	}
}

func (c *connection) exec(containerID string, interrupted chan bool) {
	defer func() {
		if err := c.execClient.Clean(containerID); err != nil {
			log.Println(err)
		}
	}()
	status, err := c.execClient.Execute(containerID, func(stdout, stderr io.Reader) {
		go c.broadcast("stdout", stdout, interrupted)
		c.broadcast("stderr", stderr, interrupted)
	})
	if err != nil {
		log.Println(err)
		return
	}
	select {
	case <-interrupted:
	default:
		c.send("status", strconv.Itoa(status))
	}
}

func (c *connection) interrupt(containerID string, interrupted chan bool) {
	for i := 0; i < 3; i++ {
		interrupted <- true
	}
	if err := c.execClient.Interrupt(containerID); err != nil {
		log.Println(err)
	}
}

func (c *connection) broadcast(stream string, output io.Reader, interrupted chan bool) {
	var (
		reader = bufio.NewReader(output)
		buffer = make([]byte, 1024)
	)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			return
		}
		select {
		case <-interrupted:
			return
		default:
			if n > 0 {
				c.send(stream, string(buffer[0:n]))
			}
		}
	}
}

func (c *connection) send(stream, chunk string) {
	response := Response{Stream: stream, Chunk: chunk}
	c.Lock()
	if err := c.ws.WriteJSON(response); err != nil {
		log.Println(err)
	}
	c.Unlock()
}
