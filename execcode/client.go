package execcode

import (
	"fmt"
	"io"

	"github.com/folieadrien/grounds/utils"
	docker "github.com/fsouza/go-dockerclient"
)

const (
	errorClientBusy           = "execcode: client is busy."
	errorClientNotBusy        = "execcode: client is not busy."
	errorLanguageNotSpecified = "execcode: language not specified."
)

type Client struct {
	docker    DockerClient
	registry  string
	container *docker.Container
	IsBusy    bool
}

func NewClient(dockerAddr, dockerRegistry string) (*Client, error) {
	docker, err := docker.NewClient(dockerAddr)
	if err != nil {
		return nil, err
	}
	return &Client{
		docker:    docker,
		registry:  dockerRegistry,
		container: nil,
		IsBusy:    false,
	}, nil
}

func (c *Client) Execute(language, code string, f func(stdout, stderr io.Reader) error) (int, error) {
	if c.IsBusy {
		return -1, fmt.Errorf(errorClientBusy)
	}
	if language == "" {
		return -1, fmt.Errorf(errorLanguageNotSpecified)
	}
	image := utils.FormatImageName(c.registry, language)
	cmd := []string{code}
	if err := c.createContainer(image, cmd); err != nil {
		return -1, err
	}

	defer c.forceRemoveContainer()

	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	go c.attachToContainer(stdoutWriter, stderrWriter)

	if err := c.docker.StartContainer(c.container.ID, c.container.HostConfig); err != nil {
		return -1, err
	}
	c.IsBusy = true
	if err := f(stdoutReader, stderrReader); err != nil {
		return -1, err
	}
	status, err := c.docker.WaitContainer(c.container.ID)
	if err != nil {
		return -1, err
	}
	return status, nil
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

func (c *Client) attachToContainer(stdout, stderr io.Writer) error {
	opts := docker.AttachToContainerOptions{
		Container:    c.container.ID,
		OutputStream: stdout,
		ErrorStream:  stderr,
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
	c.IsBusy = false
	return nil
}
