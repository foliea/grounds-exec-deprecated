package handler

import "testing"

const (
	dockerAddr       = "http://127.0.0.1:8080"
	dockerRepository = "grounds"
)

func TestNewRunHandler(t *testing.T) {
	run := NewRunHandler(true, dockerAddr, dockerRepository)
	if run == nil {
		t.Fatalf("Expected run handler to be not nil.")
	}
	if !run.upgrader.CheckOrigin(nil) {
		t.Fatalf("Expected check origin to be disabled.")
	}
}

func TestNewHandlerBadParameters(t *testing.T) {
	run := NewRunHandler(true, "", "")
	if run != nil {
		t.Fatalf("Expected run handler to be nil.")
	}
}

func TestRunHandlerServeHTTP(t *testing.T) {
	_ = newRunHandler(t)
}

func newRunHandler(t *testing.T) *RunHandler {
	return NewRunHandler(true, dockerAddr, dockerRepository)
}
