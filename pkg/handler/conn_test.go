package handler

import (
	"io"
	"testing"
)
import "github.com/gorilla/websocket"

func TestReader(t *testing.T) {
	var (
		messages = make(chan []byte)
		conn     = &connection{ws: &fakeWebsocketConn{}, receive: messages}
	)
	go conn.reader()
	for message := range messages {
		if string(message) != "test" {
			t.Fatalf("Expected to read 'test', got %v", string(message))
		}
	}
}

func TestWriter(t *testing.T) {
	var (
		messages = make(chan []byte)
		conn     = &connection{ws: &fakeWebsocketConn{}, send: messages}
	)
	go conn.writer()
	messages <- []byte("test")
}

type fakeWebsocketConn struct {
	open      bool
	sendError bool
}

func (f *fakeWebsocketConn) ReadMessage() (messageType int, p []byte, err error) {
	if f.sendError {
		return 0, nil, io.EOF
	}
	f.sendError = true
	return websocket.TextMessage, []byte("test"), nil

}

func (f *fakeWebsocketConn) WriteMessage(messageType int, data []byte) error {
	return nil
}

func (f *fakeWebsocketConn) Close() error {
	f.open = false
	return nil
}
