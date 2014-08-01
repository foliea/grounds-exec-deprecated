package execcode

import (
	"github.com/fsouza/go-dockerclient"
)

// FakeDockerClient is a simple fake docker client, so that execcode can be run for testing without requiring a real docker setup
type FakeDockerClient struct {}

func (f *FakeDockerClient) CreateContainer(c docker.CreateContainerOptions) (*docker.Container, error) {
	return &docker.Container{ID: "fake"}, nil
}

func (f *FakeDockerClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	return nil
}

func (f *FakeDockerClient) AttachToContainer(opts docker.AttachToContainerOptions) error {
	return nil
}

func (f *FakeDockerClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	return nil
}

func (f *FakeDockerClient) WaitContainer(id string) (int, error) {
	return 0, nil
}