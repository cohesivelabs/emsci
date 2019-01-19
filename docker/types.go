package docker

import docker "github.com/docker/docker/client"

type DockerClientConfig struct {
	Version string
}

type DockerClient struct {
	Daemon *docker.Client
}
