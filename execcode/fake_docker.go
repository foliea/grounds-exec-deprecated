package execcode

import (
	"errors"

	docker "github.com/fsouza/go-dockerclient"
)

var (
	ErrorContainerNotCreated = errors.New("Container not created.")
	ErrorInvalidImage        = errors.New("Invalid image.")
	ErrorInvalidOpts         = errors.New("Invalid opts.")
)

// FakeDockerClient is a simple fake docker client, so that execcode can be run for testing without requiring a real docker setup
type FakeDockerClient struct {
	container *docker.Container
}

func (f *FakeDockerClient) CreateContainer(c docker.CreateContainerOptions) (*docker.Container, error) {
	if c.Config.Image == "" {
		return nil, ErrorInvalidImage
	}
	if !c.Config.AttachStdout || !c.Config.AttachStderr || !c.Config.NetworkDisabled {
		return nil, ErrorInvalidOpts
	}
	f.container = &docker.Container{ID: "fake"}
	return f.container, nil
}

func (f *FakeDockerClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	if f.container == nil {
		return ErrorContainerNotCreated
	}
	return nil
}

func (f *FakeDockerClient) AttachToContainer(opts docker.AttachToContainerOptions) error {
	if f.container == nil {
		return ErrorContainerNotCreated
	}
	if opts.Container == "" || opts.OutputStream == nil || opts.ErrorStream == nil ||
		!opts.Stream || !opts.Stdout || !opts.Stderr {
		return ErrorInvalidOpts
	}
	return nil
}

func (f *FakeDockerClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	if f.container == nil {
		return ErrorContainerNotCreated
	}
	if opts.ID == "" || opts.Force == false {
		return ErrorInvalidOpts
	}
	f.container = nil
	return nil
}

func (f *FakeDockerClient) WaitContainer(id string) (int, error) {
	if f.container == nil {
		return -1, ErrorContainerNotCreated
	}
	return 0, nil
}
