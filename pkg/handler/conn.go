package handler

import (
	"log"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

var requestDelay = 500 * time.Millisecond

type Connection struct {
	input  chan []byte
	output chan []byte
	time   time.Time
	event  string
	so     socketio.Socket
}

func (c *Connection) Read(msg string) {
	// There is a bug with go-socket-io, this is a trick to prevent it before it gets patched upstream
	c.so.Emit("msg", msg)

	log.Println("received: ", msg)
	requestTime := time.Now()
	if requestTime.Sub(c.time) <= requestDelay {
		log.Println("ignored: ", msg)
		return
	}
	c.time = requestTime
	c.input <- []byte(msg)
}

func (c *Connection) Write() {
	for msg := range c.output {
		response := string(msg[0:len(msg)])

		log.Println("sent: ", response)
		c.so.Emit(c.event, response)
	}
}
