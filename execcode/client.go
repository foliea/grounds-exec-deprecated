package execcode

import (
	"io"
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

const (
	errorClientBusy = "execcode: client is busy"
	errorClientNotBusy = "execcode: client is not busy"
)

type Client struct {
	docker DockerInterface
	registry string
	container *docker.Container
	IsBusy bool
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
		IsBusy: false,
	}, nil
}

func (c *Client) Execute(language, code string, f func(stdout, stderr io.Reader) error) error {
	if c.IsBusy {
		return fmt.Errorf(errorClientBusy)
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

	if err := c.docker.StartContainer(c.container.ID, c.container.HostConfig); err != nil {
		return err
	}
	c.IsBusy = true
	f(stdoutReader, stderrReader) // FIXME: Handle f error
	if _, err := c.docker.WaitContainer(c.container.ID); err != nil {
		return err
	}
	// FIXME: do something with status
	return nil
}

func (c *Client) Interrupt() error {
	if !c.IsBusy {
		return fmt.Errorf(errorClientNotBusy)
	}
	if err := c.forceRemoveContainer(); err != nil {
		return err
	}
	return nil
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

func (c *Client) forceRemoveContainer() error {
	opts := docker.RemoveContainerOptions{
		ID: c.container.ID,
		Force: true,
	}
	if err := c.docker.RemoveContainer(opts); err != nil {
		return err
	}
	c.IsBusy = false
	return nil
}

