package execcode

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
		t.Errorf("Expected registry %s. Got %s.", registry, client.registry)
	}
}

func TestNewClientInvalidEndpoint(t *testing.T) {
	_, err := NewClient(invalidEndpoint, "")
	if err == nil {
		t.Errorf("Expected error. Got nothing.")
	}
}

func TestPrepare(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	if containerID == "" {
		t.Fatalf("Expected containerID. Got empty string.")
	}
}

func TestPrepareWithEmptyLanguage(t *testing.T) {
	client := newFakeClient(t)
	containerID, err := client.Prepare("", "puts 42")
	if err == nil {
		t.Fatalf("Expected error. Got nothing.")
	}
	if err != ErrorLanguageNotSpecified {
		t.Fatalf("Expected error to be %v. Got %v.", ErrorLanguageNotSpecified, err)
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
		t.Fatalf("Expected attach to be true. Got false.")
	}
	if status != 0 {
		t.Fatalf("Expected status to be 0. Got %v.", status)
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
	client.docker = &FakeDockerClient{}
	return client
}
