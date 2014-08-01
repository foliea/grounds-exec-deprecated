package execcode

import "github.com/fsouza/go-dockerclient"

// DockerInterface is an abstract interface for testability. It abstracts the interface of docker.Client
type DockerInterface interface {
	CreateContainer(docker.CreateContainerOptions) (*docker.Container, error)
	StartContainer(id string, hostConfig *docker.HostConfig) error
	AttachToContainer(opts docker.AttachToContainerOptions) error
	RemoveContainer(opts docker.RemoveContainerOptions) error
	WaitContainer(id string) (int, error)
}