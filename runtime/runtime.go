package runtime

import (
	"emsci/runtime/docker"
)

func NewDockerClient() (docker.DockerClient, error) {
	return docker.NewClient()
}
