package execcode

import (
	"io"
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

const (
	ErrorClientBusy = "execcode: client is busy"
	ErrorClientNotBusy = "execcode: client is not busy"
)

type Client struct {
	docker *docker.Client
	registry string
	container *docker.Container
}

func NewClient(dockerAddr, dockerRegistry string) (*Client, error) {
	docker, err := docker.NewClient(dockerAddr)
	if (err != nil) {
		return nil, err
	}
	return &Client{
		docker: docker,
		registry: dockerRegistry,
		container: nil,
	}, nil
}

func (c *Client) Execute(language, code string, f func(stdout, stderr io.Reader) error) error {
	if c.IsBusy() {
		return fmt.Errorf(ErrorClientBusy)
	}
	image := fmt.Sprintf("%s/exec-%s", c.registry, language)
	cmd := []string{code}

	if err := c.createContainer(image, cmd); err != nil {
		return err
	}

	defer c.forceRemoveContainer()

	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	go c.attachToContainer(stdoutWriter, stderrWriter)

	if err := c.startContainer(); err !=nil {
		return err
	}
	return f(stdoutReader, stderrReader)
}

func (c *Client) Interrupt() error {
	if !c.IsBusy() {
		return fmt.Errorf(ErrorClientNotBusy)
	}
	if err := c.forceRemoveContainer(); err != nil {
		return err
	}
	return nil
}

func (c *Client) IsBusy() bool {
	if c.container == nil {
		return false
	}
	return true
}

func (c *Client) createContainer(image string, cmd []string) error {
	config := &docker.Config{
		Cmd: cmd,
		Image: image,
		AttachStdout: true,
		AttachStderr: true,
		NetworkDisabled: true,
	}
	opts := docker.CreateContainerOptions{
		Name: "",
		Config: config,
	}
	var err error
	if c.container, err = c.docker.CreateContainer(opts); err != nil {
		return err
	}
	return nil
}

func (c *Client) attachToContainer(stdout, stderr io.Writer) error {
	opts := docker.AttachToContainerOptions{
		Container: c.container.ID,
		OutputStream: stdout,
		ErrorStream: stderr,
		Stream: true,
		Stdin: true,
		Stdout: true,
	}
	if err := c.docker.AttachToContainer(opts); err != nil {
		return err
	}
	return nil
}

func (c *Client) startContainer() error {
	if err := c.docker.StartContainer(c.container.ID, c.container.HostConfig); err != nil {
		return err
	}
	return nil
}

func (c *Client) forceRemoveContainer() error {
	opts := docker.RemoveContainerOptions{
		ID: c.container.ID,
		Force: true,
	}
	if err := c.docker.RemoveContainer(opts); err != nil {
		return err
	}
	c.container = nil
	return nil
}

