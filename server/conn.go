package main

import (
	"bufio"
	"fmt"
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

func (c *connection) read() {
	defer c.ws.Close()
	containerID := ""
	var interrupted chan bool
	for {
		input := Input{}
		// Read input from client
		err := c.ws.ReadJSON(&input)
		// Interrupt previous container execution
		if containerID != "" {
			go c.interrupt(containerID, interrupted)
		}
		//FIXME: timeout excess
		if err != nil {
			log.Println(err)
			return
		}

		// FIXME: IF CODE > 1500 program too large
		// Prepare a new container
		if containerID, err = c.execClient.Prepare(input.Language, input.Code); err != nil {
			log.Println(err)
			continue
		}
		interrupted = make(chan bool, 3)

		// Execute code inside the new container
		go c.exec(containerID, interrupted)
	}
}

func (c *connection) exec(containerID string, interrupted chan bool) {
	// Execute code with execcode and send output to the client
	status, err := c.execClient.Execute(containerID, func(stdout, stderr io.Reader) {
		go c.broadcast("stdout", stdout, interrupted)
		c.broadcast("stderr", stderr, interrupted)
	})
	select {
	case <-interrupted:
	default:
		if err != nil {
			log.Println(err)
		} else {
			c.send("status", strconv.Itoa(status))
		}
	}
	// Cleanup execcode container
	fmt.Println("Cleaning: ", containerID)
	if err := c.execClient.Clean(containerID); err != nil {
		log.Println(err)
	}
}

func (c *connection) interrupt(containerID string, interrupted chan bool) {
	interrupted <- true
	interrupted <- true
	interrupted <- true
	if err := c.execClient.Interrupt(containerID); err != nil {
		log.Println(err)
	}
}

func (c *connection) broadcast(stream string, output io.Reader, interrupted chan bool) {
	reader := bufio.NewReader(output)
	buffer := make([]byte, 1024) // check other projects and verify read size
	// with big code example > 500 lines
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
