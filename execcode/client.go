package execcode

import (
	"errors"
	"io"
	"sync"

	"github.com/folieadrien/grounds/utils"
	docker "github.com/fsouza/go-dockerclient"
)

var (
	ErrorClientBusy           = errors.New("execcode: client is busy.")
	ErrorClientNotBusy        = errors.New("execcode: client is not busy.")
	ErrorLanguageNotSpecified = errors.New("execcode: language not specified.")
)

type Client struct {
	docker      DockerClient
	registry    string
	container   *docker.Container
	mu          sync.Mutex
	stdout      *io.PipeWriter
	stderr      *io.PipeWriter
	interrupted bool
	IsBusy      bool
}

func NewClient(dockerAddr, dockerRegistry string) (*Client, error) {
	docker, err := docker.NewClient(dockerAddr)
	if err != nil {
		return nil, err
	}
	return &Client{
		docker:      docker,
		registry:    dockerRegistry,
		container:   nil,
		interrupted: false,
		IsBusy:      false,
	}, nil
}

func (c *Client) Execute(language, code string, f func(stdout, stderr io.Reader)) (int, error) {
	if c.IsBusy {
		return -1, ErrorClientBusy
	}
	if language == "" {
		return -1, ErrorLanguageNotSpecified
	}
	image := utils.FormatImageName(c.registry, language)
	cmd := []string{utils.FormatCode(code)}
	if err := c.createContainer(image, cmd); err != nil {
		return -1, err
	}

	var stdoutReader, stderrReader *io.PipeReader
	stdoutReader, c.stdout = io.Pipe()
	stderrReader, c.stderr = io.Pipe()

	defer c.closeRessources()

	c.IsBusy = true
	c.setInterrupted(false)

	go c.attachToContainer()

	if err := c.docker.StartContainer(c.container.ID, c.container.HostConfig); err != nil {
		return -1, err
	}

	go f(stdoutReader, stderrReader)

	status, err := c.docker.WaitContainer(c.container.ID)
	if err != nil {
		return -1, err
	}
	return status, nil
}

func (c *Client) Interrupt() error {
	if !c.IsBusy {
		return ErrorClientNotBusy
	}
	c.setInterrupted(true)
	if err := c.closeRessources(); err != nil {
		return err
	}
	return nil
}

func (c *Client) Interrupted() bool {
	c.mu.Lock()
	value := c.interrupted
	c.mu.Unlock()
	return value
}

func (c *Client) setInterrupted(value bool) {
	c.mu.Lock()
	c.interrupted = value
	c.mu.Unlock()
}

func (c *Client) closeRessources() error {
	c.IsBusy = false
	c.stdout.Close()
	c.stderr.Close()
	if err := c.forceRemoveContainer(); err != nil {
		return err
	}
	return nil
}

func (c *Client) createContainer(image string, cmd []string) error {
	config := &docker.Config{
		Cmd:             cmd,
		Image:           image,
		AttachStdout:    true,
		AttachStderr:    true,
		NetworkDisabled: true,
	}
	opts := docker.CreateContainerOptions{
		Name:   "",
		Config: config,
	}
	var err error
	if c.container, err = c.docker.CreateContainer(opts); err != nil {
		return err
	}
	return nil
}

func (c *Client) attachToContainer() error {
	opts := docker.AttachToContainerOptions{
		Container:    c.container.ID,
		OutputStream: c.stdout,
		ErrorStream:  c.stderr,
		Stream:       true,
		Stdout:       true,
		Stderr:       true,
	}
	if err := c.docker.AttachToContainer(opts); err != nil {
		return err
	}
	return nil
}

func (c *Client) forceRemoveContainer() error {
	opts := docker.RemoveContainerOptions{
		ID:    c.container.ID,
		Force: true,
	}
	if err := c.docker.RemoveContainer(opts); err != nil {
		return err
	}
	return nil
}
