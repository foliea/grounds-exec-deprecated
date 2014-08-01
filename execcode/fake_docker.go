package execcode

import (
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

const (
	errorContainerNotCreated = "Container not created."
	errorImageInvalid = "Image invalid."
) 
	

// FakeDockerClient is a simple fake docker client, so that execcode can be run for testing without requiring a real docker setup
type FakeDockerClient struct {
	container *docker.Container
}

func (f *FakeDockerClient) CreateContainer(c docker.CreateContainerOptions) (*docker.Container, error) {
	if c.Config.Image == "" {
		return nil, fmt.Errorf(errorImageInvalid)
	}
	f.container = &docker.Container{ID: "fake"}
	return  f.container, nil
}

func (f *FakeDockerClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	if f.container == nil {
		return fmt.Errorf(errorContainerNotCreated)
	}
	return nil
}

func (f *FakeDockerClient) AttachToContainer(opts docker.AttachToContainerOptions) error {
	if f.container == nil {
		return fmt.Errorf(errorContainerNotCreated)
	}
	return nil
}

func (f *FakeDockerClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	if f.container == nil {
		return fmt.Errorf(errorContainerNotCreated)
	}
	f.container = nil
	return nil
}

func (f *FakeDockerClient) WaitContainer(id string) (int, error) {
	if f.container == nil {
		return -1, fmt.Errorf(errorContainerNotCreated)
	}
	return 0, nil
}