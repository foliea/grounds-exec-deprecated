package main

import (
	"testing"

	"github.com/folieadrien/grounds/execcode"
)

func TestNewRunHandler(t *testing.T) {
	debug := true
	execClient := newExecClient(t)
	run := NewRunHandler(debug, execClient)
	if run == nil {
		t.Fatalf("Expected run handler to be not nil.")
	}
	if !run.upgrader.CheckOrigin(nil) {
		t.Fatalf("Expected check origin to be disabled.")
	}
}

func TestNewRunHandlerWithBadParameters(t *testing.T) {
	run := NewRunHandler(true, nil)
	if run != nil {
		t.Fatalf("Expected run handler to be nil.")
	}
}

func TestRunHandlerServeHTTP(t *testing.T) {
	_ = newRunHandler(t)
}

func newRunHandler(t *testing.T) *RunHandler {
	return NewRunHandler(true, newExecClient(t))
}

func newExecClient(t *testing.T) *execcode.Client {
	execClient, err := execcode.NewClient("http://127.0.0.1:8080", "grounds")
	if err != nil {
		t.Fatal(err)
	}
	return execClient
}
