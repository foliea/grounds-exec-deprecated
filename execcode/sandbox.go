package execcode

import (
	"log"

	"github.com/samalba/dockerclient"
)

func Execute() {
	docker, _ := dockerclient.NewDockerClient("http://192.168.59.103:2375", nil)

	containers, err := docker.ListContainers(true)
	if err != nil {
		log.Fatal(err)
	}
	for _, c := range containers {
		log.Println(c.Id, c.Names)
	}
}
