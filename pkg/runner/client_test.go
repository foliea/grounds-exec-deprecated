package runner

import (
	"bufio"
	"io"
	"testing"
)

const (
	validEndpoint   = "http://localhost:4243"
	invalidEndpoint = ""
)

func TestNewClient(t *testing.T) {
	registry := "test"
	client, err := NewClient(validEndpoint, registry)
	if err != nil {
		t.Fatal(err)
	}
	if client.registry != registry {
		t.Errorf("Expected registry %s, got %s.", registry, client.registry)
	}
}

func TestNewClientInvalidEndpoint(t *testing.T) {
	_, err := NewClient(invalidEndpoint, "")
	if err == nil {
		t.Errorf("Expected error, got nothing.")
	}
}

func TestPrepare(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	if containerID == "" {
		t.Fatalf("Expected containerID, got empty string.")
	}
}

func TestPrepareProgrameToolarge(t *testing.T) {
	var (
		client = newFakeClient(t)
		code   = getTooLargeProgram()
	)
	containerID, err := client.Prepare("ruby", code)
	if err == nil {
		t.Fatalf("Expected error, got nothing.")
	}
	if err != ErrorProgramTooLarge {
		t.Fatalf("Expected error to be %v, got %v.", ErrorProgramTooLarge, err)
	}
	if containerID != "" {
		t.Fatalf("Expected containerID to be empty.")
	}
}

func TestPrepareCreateFailed(t *testing.T) {
	client := newFakeClient(t)
	client.docker = NewFakeDockerClient(&FakeDockerClientOptions{createFail: true})
	containerID, err := client.Prepare("ruby", "puts 42")
	if err == nil {
		t.Fatalf("Expected error, got nothing.")
	}
	if containerID != "" {
		t.Fatalf("Expected containerID to be empty.")
	}
}

func TestPrepareWithEmptyLanguage(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("", "puts 42")
	if err == nil {
		t.Fatalf("Expected error, got nothing.")
	}
	if err != ErrorLanguageNotSpecified {
		t.Fatalf("Expected error to be %v, got %v.", ErrorLanguageNotSpecified, err)
	}
	if containerID != "" {
		t.Fatalf("Expected containerID to be empty.")
	}
}

func TestExecute(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	attach := false
	status, err := client.Execute(containerID, func(stdout, stderr io.Reader) {
		attach = true

		readOut := bufio.NewReader(stdout)
		if _, err := readOut.Read(make([]byte, 2)); err != io.EOF {
			t.Fatalf("Expected stdout to be closed.")
		}
		readErr := bufio.NewReader(stderr)
		if _, err := readErr.Read(make([]byte, 2)); err != io.EOF {
			t.Fatalf("Expected stderr to be closed.")
		}
	})
	if err != nil {
		t.Fatal(err)
	}
	if attach == false {
		t.Fatalf("Expected attach to be true, got false.")
	}
	if status != 0 {
		t.Fatalf("Expected status to be 0, got %v.", status)
	}
}

func TestExecuteNotPrepared(t *testing.T) {
	client := newFakeClient(t)
	attach := false
	_, err := client.Execute("-1", func(stdout, stderr io.Reader) {
		attach = true
	})
	if err == nil {
		t.Fatal("Expected an error, got nothing.")
	}
	if attach == true {
		t.Fatalf("Expected attach to be false, got true.")
	}
}

func TestExecuteAttachFailed(t *testing.T) {
	client := newFakeClient(t)
	client.docker = NewFakeDockerClient(&FakeDockerClientOptions{attachFail: true})
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Execute(containerID, func(stdout, stderr io.Reader) {})
	if err == nil {
		t.Fatal("Expected an error, got nothing.")
	}
	if err != ErrorAttachFailed {
		t.Fatalf("Expected error to be %v, got %v.", ErrorAttachFailed, err)
	}
}

func TestExecuteWaitFailed(t *testing.T) {
	client := newFakeClient(t)
	client.docker = NewFakeDockerClient(&FakeDockerClientOptions{waitFail: true})
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	attach := false
	_, err = client.Execute(containerID, func(stdout, stderr io.Reader) {
		attach = true
	})
	if err == nil {
		t.Fatal("Expected an error, got nothing.")
	}
	if err != ErrorWaitFailed {
		t.Fatalf("Expected error to be %v, got %v.", ErrorWaitFailed, err)
	}
	if attach == true {
		t.Fatalf("Expected attach to be false, got true.")
	}
}

func TestInterrupt(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	_, err = client.Execute(containerID, func(stdout, stderr io.Reader) {})
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Interrupt(containerID); err != nil {
		t.Fatal(err)
	}
}

func TestClean(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Clean(containerID); err != nil {
		t.Fatal(err)
	}
}

func newFakeClient(t *testing.T) *Client {
	client, err := NewClient(validEndpoint, "")
	if err != nil {
		t.Fatal(err)
	}
	client.docker = NewFakeDockerClient(nil)
	return client
}

func getTooLargeProgram() string {
	code := make([]byte, programMaxSize)
	return string(code[0:programMaxSize])
}
