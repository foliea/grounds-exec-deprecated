package execcode

import (
	"errors"
	"io"

	"github.com/folieadrien/grounds/utils"
	docker "github.com/fsouza/go-dockerclient"
)

const programMaxSize = 65536

var (
	ErrorLanguageNotSpecified = errors.New("Language not specified.")
	ErrorProgramTooLarge      = errors.New("Program too large.")
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
	if len(code) >= programMaxSize {
		return "", ErrorProgramTooLarge
	}
	var (
		image = utils.FormatImageName(c.registry, language)
		cmd   = []string{utils.FormatCode(code)}
	)
	container, err := c.createContainer(image, cmd)
	if err != nil {
		return "", err
	}
	return container.ID, nil
}

func (c *Client) Execute(containerID string, attach func(stdout, stderr io.Reader)) (int, error) {
	var (
		stdoutReader, stdoutWriter = io.Pipe()
		stderrReader, stderrWriter = io.Pipe()
	)
	defer func() {
		stdoutWriter.Close()
		stderrWriter.Close()
	}()

	errs := make(chan error, 1)
	go func() {
		errs <- c.attachToContainer(containerID, stdoutWriter, stderrWriter)
	}()
	go attach(stdoutReader, stderrReader)

	if err := c.docker.StartContainer(containerID, nil); err != nil {
		return 0, err
	}
	status, err := c.docker.WaitContainer(containerID)
	if err != nil {
		return 0, err
	}
	if err := <-errs; err != nil {
		return 0, err
	}
	return status, nil
}

func (c *Client) Clean(containerID string) error {
	return c.removeContainer(containerID)
}

func (c *Client) Interrupt(containerID string) error {
	return c.docker.StopContainer(containerID, 5)
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
