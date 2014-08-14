package main

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"sync"

	"github.com/folieadrien/grounds/pkg/execcode"
	"github.com/gorilla/websocket"
)

type connection struct {
	writeLock  sync.Mutex
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
			return
		}
		if containerID, err = c.execClient.Prepare(input.Language, input.Code); err != nil {
			c.handleError(err)
			continue
		}
		interrupted = make(chan bool, 3)
		go c.exec(containerID, interrupted)
	}
}

func (c *connection) exec(containerID string, interrupted chan bool) {
	defer func() {
		if err := c.execClient.Clean(containerID); err != nil {
			c.handleError(err)
		}
	}()
	status, err := c.execClient.Execute(containerID, func(stdout, stderr io.Reader) {
		go c.broadcast("stdout", stdout, interrupted)
		c.broadcast("stderr", stderr, interrupted)
	})
	if err != nil {
		c.handleError(err)
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
	c.execClient.Interrupt(containerID)
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
	c.writeLock.Lock()
	if err := c.ws.WriteJSON(response); err != nil {
		log.Println(err)
	}
	c.writeLock.Unlock()
}

// Send the error to the client or log the error server side
func (c *connection) handleError(err error) {
	if err == execcode.ErrorLanguageNotSpecified ||
		err == execcode.ErrorProgramTooLarge {
		c.send("error", err.Error())
	} else {
		log.Println(err)
	}
}
