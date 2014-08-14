package execcode

import (
	"errors"

	docker "github.com/fsouza/go-dockerclient"
)

var (
	ErrorNoSuchContainerID   = errors.New("No such container ID.")
	ErrorContainerNotStarted = errors.New("Container not started.")
	ErrorInvalidImage        = errors.New("Invalid image.")
	ErrorCreateInvalidOpts   = errors.New("Create invalid opts.")
	ErrorAttachInvalidOpts   = errors.New("Attach invalid opts.")
	ErrorRemoveInvalidOpts   = errors.New("Remove invalid opts.")
	ErrorAttachFailed        = errors.New("Attach failed.")
	ErrorCreateFailed        = errors.New("Create failed.")
	ErrorWaitFailed          = errors.New("Wait failed.")
)

// FakeDockerClient is a simple fake docker client, so that execcode can be run for testing without requiring a real docker setup
type FakeDockerClient struct {
	container        *docker.Container
	containerStarted bool
	opts             *FakeDockerClientOptions
}

type FakeDockerClientOptions struct {
	createFail bool
	attachFail bool
	waitFail   bool
}

func NewFakeDockerClient(opts *FakeDockerClientOptions) dockerClient {
	if opts == nil {
		opts = &FakeDockerClientOptions{}
	}
	return &FakeDockerClient{container: &docker.Container{}, opts: opts}
}

func (f *FakeDockerClient) CreateContainer(c docker.CreateContainerOptions) (*docker.Container, error) {
	if f.opts.createFail {
		return nil, ErrorCreateFailed
	}
	if c.Config.Image == "" {
		return nil, ErrorInvalidImage
	}
	if !c.Config.AttachStdout || !c.Config.AttachStderr || !c.Config.NetworkDisabled {
		return nil, ErrorCreateInvalidOpts
	}
	f.container = &docker.Container{ID: "fake"}
	f.containerStarted = false
	return f.container, nil
}

func (f *FakeDockerClient) AttachToContainer(opts docker.AttachToContainerOptions) error {
	if f.opts.attachFail {
		return ErrorAttachFailed
	}
	if f.container.ID != opts.Container {
		return ErrorNoSuchContainerID
	}
	if opts.Container == "" || opts.OutputStream == nil || opts.ErrorStream == nil ||
		!opts.Stream || !opts.Stdout || !opts.Stderr {
		return ErrorAttachInvalidOpts
	}
	return nil
}

func (f *FakeDockerClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	if f.container.ID != id {
		return ErrorNoSuchContainerID
	}
	f.containerStarted = true
	return nil
}

func (f *FakeDockerClient) WaitContainer(id string) (int, error) {
	if f.container.ID != id {
		return -1, ErrorNoSuchContainerID
	}
	if f.containerStarted == false {
		return -1, ErrorContainerNotStarted
	}
	if f.opts.waitFail {
		return -1, ErrorWaitFailed
	}
	return 0, nil
}

func (f *FakeDockerClient) StopContainer(id string, timeout uint) error {
	if f.container.ID != id {
		return ErrorNoSuchContainerID
	}
	if f.containerStarted == false {
		return ErrorContainerNotStarted
	}
	f.containerStarted = false
	return nil
}

func (f *FakeDockerClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	if f.container.ID != opts.ID {
		return ErrorNoSuchContainerID
	}
	if opts.ID == "" || opts.Force == true {
		return ErrorRemoveInvalidOpts
	}
	f.container = nil
	return nil
}
