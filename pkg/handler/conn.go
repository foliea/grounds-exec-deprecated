package handler

import "github.com/gorilla/websocket"

type websocketConn interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	Close() error
}

type connection struct {
	ws websocketConn

	// Buffered channel of inbound messages.
	receive chan []byte
	// Buffered channel of outbound messages.
	send chan []byte
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		c.receive <- message
	}
	close(c.receive)
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		if err := c.ws.WriteMessage(websocket.TextMessage, message); err != nil {
			break
		}
	}
	c.ws.Close()
}
