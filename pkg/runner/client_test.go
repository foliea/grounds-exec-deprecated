package runner

import (
	"bufio"
	"io"
	"testing"
	"time"

	docker "github.com/fsouza/go-dockerclient"
	dockerTest "github.com/fsouza/go-dockerclient/testing"
)

func TestNewClient(t *testing.T) {
	repository := "test"
	client, err := NewClient("http://localhost:4243", repository)
	if err != nil {
		t.Fatal(err)
	}
	if client.repository != repository {
		t.Errorf("Expected registry %s, got %s.", repository, client.repository)
	}
}

func TestNewClientInvalidEndpoint(t *testing.T) {
	_, err := NewClient("", "")
	if err == nil {
		t.Errorf("Expected error, got nothing.")
	}
}

func TestPrepare(t *testing.T) {
	client, server := NewFakeClientServer(t, nil)
	defer server.Stop()

	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	if containerID == "" {
		t.Fatalf("Expected containerID, got empty string.")
	}
}

func TestPrepareProgrameToolarge(t *testing.T) {
	client, server := NewFakeClientServer(t, nil)
	defer server.Stop()

	containerID, err := client.Prepare("ruby", getTooLargeProgram())
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

func TestPrepareWithEmptyLanguage(t *testing.T) {
	client, server := NewFakeClientServer(t, nil)
	defer server.Stop()

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
	cChan := make(chan *docker.Container, 2)
	client, server := NewFakeClientServer(t, cChan)
	defer server.Stop()

	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	// Wait that the container is properly created
	<-cChan
	end := make(chan bool)
	go func() {
		status, err := client.Execute(containerID, func(stdout, stderr io.Reader) {
			var (
				buff    = make([]byte, 1024)
				readOut = bufio.NewReader(stdout)
				readErr = bufio.NewReader(stderr)
			)
			for {
				if _, err := readOut.Read(buff); err != nil {
					break
				}
			}
			for {
				if _, err := readErr.Read(buff); err != nil {
					break
				}
			}
			end <- true
		})
		if err != nil {
			t.Fatal(err)
		}
		if status != 0 {
			t.Fatalf("Expected status to be 0, got %v.", status)
		}
	}()
	// Wait that the container is properly started
	<-cChan
	if err := client.Interrupt(containerID); err != nil {
		t.Fatal(err)
	}
	select {
	case <-end:
	case <-time.After(time.Second):
		t.Fatalf("Expected stdout and stderr to be closed.")
	}
}

func TestExecuteNotPrepared(t *testing.T) {
	client, server := NewFakeClientServer(t, nil)
	defer server.Stop()

	_, err := client.Execute("-1", func(stdout, stderr io.Reader) {})
	if err == nil {
		t.Fatal("Expected an error, got nothing.")
	}
}

func TestInterrupt(t *testing.T) {
	cChan := make(chan *docker.Container, 2)
	client, server := NewFakeClientServer(t, cChan)
	defer server.Stop()

	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	// Wait that the container is properly created
	<-cChan
	end := make(chan bool)
	go func() {
		_, err := client.Execute(containerID, func(stdout, stderr io.Reader) {})
		if err != nil {
			t.Fatal(err)
		}
		end <- true
	}()
	// Wait that the container is properly started
	<-cChan
	if err := client.Interrupt(containerID); err != nil {
		t.Fatal(err)
	}
	select {
	case <-end:
	case <-time.After(time.Second):
		t.Fatalf("Expected execution to be interrupted.")
	}
}

func TestClean(t *testing.T) {
	client, server := NewFakeClientServer(t, nil)
	defer server.Stop()

	containerID, err := client.Prepare("ruby", "puts 42")
	if err != nil {
		t.Fatal(err)
	}
	if err := client.Clean(containerID); err != nil {
		t.Fatal(err)
	}
	if err := client.Interrupt(containerID); err == nil {
		t.Fatal("Expected container %v to be removed.", containerID)
	}
}

func NewFakeClientServer(t *testing.T, cChan chan *docker.Container) (*Client, *dockerTest.DockerServer) {
	server, err := dockerTest.NewServer("127.0.0.1:0", cChan, nil)
	if err != nil {
		t.Fatal(err)
	}
	client, err := NewClient(server.URL(), "grounds")
	if err != nil {
		server.Stop()
		t.Fatal(err)
	}
	var (
		opts = docker.PullImageOptions{Repository: "grounds/exec-ruby"}
		auth = docker.AuthConfiguration{}
	)
	if err := client.docker.PullImage(opts, auth); err != nil {
		server.Stop()
		t.Fatal(err)
	}
	return client, server
}

func getTooLargeProgram() string {
	code := make([]byte, programMaxSize)
	return string(code[0:programMaxSize])
}
