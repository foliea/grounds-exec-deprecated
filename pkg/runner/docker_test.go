package runner

import (
	"errors"

	docker "github.com/fsouza/go-dockerclient"
)

var (
	errorNoSuchContainerID   = errors.New("No such container ID.")
	errorContainerNotStarted = errors.New("Container not started.")
	errorInvalidImage        = errors.New("Invalid image.")
	errorCreateInvalidOpts   = errors.New("Create invalid opts.")
	errorAttachInvalidOpts   = errors.New("Attach invalid opts.")
	errorRemoveInvalidOpts   = errors.New("Remove invalid opts.")
	errorAttachFailed        = errors.New("Attach failed.")
	errorCreateFailed        = errors.New("Create failed.")
	errorWaitFailed          = errors.New("Wait failed.")
)

// FakeDockerClient is a simple fake docker client, so that runner can be run for testing without requiring a real docker setup
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
		return nil, errorCreateFailed
	}
	if c.Config.Image == "" {
		return nil, errorInvalidImage
	}
	if !c.Config.AttachStdout || !c.Config.AttachStderr || !c.Config.NetworkDisabled {
		return nil, errorCreateInvalidOpts
	}
	f.container = &docker.Container{ID: "fake"}
	f.containerStarted = false
	return f.container, nil
}

func (f *FakeDockerClient) AttachToContainer(opts docker.AttachToContainerOptions) error {
	if f.opts.attachFail {
		return errorAttachFailed
	}
	if f.container.ID != opts.Container {
		return errorNoSuchContainerID
	}
	if opts.Container == "" || opts.OutputStream == nil || opts.ErrorStream == nil ||
		!opts.Stream || !opts.Stdout || !opts.Stderr {
		return errorAttachInvalidOpts
	}
	return nil
}

func (f *FakeDockerClient) StartContainer(id string, hostConfig *docker.HostConfig) error {
	if f.container.ID != id {
		return errorNoSuchContainerID
	}
	f.containerStarted = true
	return nil
}

func (f *FakeDockerClient) WaitContainer(id string) (int, error) {
	if f.container.ID != id {
		return -1, errorNoSuchContainerID
	}
	if f.containerStarted == false {
		return -1, errorContainerNotStarted
	}
	if f.opts.waitFail {
		return -1, errorWaitFailed
	}
	return 0, nil
}

func (f *FakeDockerClient) StopContainer(id string, timeout uint) error {
	if f.container.ID != id {
		return errorNoSuchContainerID
	}
	if f.containerStarted == false {
		return errorContainerNotStarted
	}
	f.containerStarted = false
	return nil
}

func (f *FakeDockerClient) RemoveContainer(opts docker.RemoveContainerOptions) error {
	if f.container.ID != opts.ID {
		return errorNoSuchContainerID
	}
	if opts.ID == "" || opts.Force == true {
		return errorRemoveInvalidOpts
	}
	f.container = nil
	return nil
}
