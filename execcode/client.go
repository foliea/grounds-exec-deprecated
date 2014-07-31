package execcode

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

const (
	ErrorClientBusy = "execcode: client is busy"
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

func (c *Client) Execute(language, code string, f func() error) error {
	if c.IsBusy() {
		return fmt.Errorf(ErrorClientBusy)
	}
	image := fmt.Sprintf("%s/exec-%s", c.registry, language)
	cmd := []string{code}
	if err := c.createContainer(image, cmd); err != nil {
		return err
	}
	return nil
}

func (c *Client) Interrupt() error {
	return nil
}

func (c *Client) IsBusy() bool {
	return false
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
