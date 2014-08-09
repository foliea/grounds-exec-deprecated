package execcode

import (
	"errors"
	"io"

	"github.com/folieadrien/grounds/utils"
	docker "github.com/fsouza/go-dockerclient"
)

var (
	ErrorLanguageNotSpecified = errors.New("execcode: language not specified.")
)

type Client struct {
	docker   dockerClient
	registry string
}

func NewClient(dockerAddr, dockerRegistry string) (*Client, error) {
	docker, err := docker.NewClient(dockerAddr)
	if err != nil {
		return nil, err
	}
	return &Client{
		docker:   docker,
		registry: dockerRegistry,
	}, nil
}

func (c *Client) Prepare(language, code string) (string, error) {
	if language == "" {
		return "", ErrorLanguageNotSpecified
	}
	image := utils.FormatImageName(c.registry, language)
	cmd := []string{utils.FormatCode(code)}

	container, err := c.createContainer(image, cmd)
	if err != nil {
		return "", err
	}
	return container.ID, nil
}

func (c *Client) Execute(containerID string, attach func(stdout, stderr io.Reader)) error {
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	defer func() {
		stdoutWriter.Close()
		stderrWriter.Close()
	}()

	go c.attachToContainer(containerID, stdoutWriter, stderrWriter)
	go attach(stdoutReader, stderrReader)

	if err := c.docker.StartContainer(containerID, nil); err != nil {
		return err
	}
	if _, err := c.docker.WaitContainer(containerID); err != nil {
		return err
	}
	return nil
}

func (c *Client) Clean(containerID string) error {
	return c.removeContainer(containerID)
}

func (c *Client) Interrupt(containerID string) error {
	return c.docker.StopContainer(containerID, 0)
}

func (c *Client) createContainer(image string, cmd []string) (*docker.Container, error) {
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
	return c.docker.CreateContainer(opts)
}

func (c *Client) attachToContainer(containerID string, stdout, stderr io.Writer) error {
	opts := docker.AttachToContainerOptions{
		Container:    containerID,
		OutputStream: stdout,
		ErrorStream:  stderr,
		Stream:       true,
		Stdout:       true,
		Stderr:       true,
	}
	return c.docker.AttachToContainer(opts)
}

func (c *Client) removeContainer(containerID string) error {
	opts := docker.RemoveContainerOptions{
		ID:    containerID,
		Force: false,
	}
	return c.docker.RemoveContainer(opts)
}
