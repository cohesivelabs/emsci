package docker

import (
	docker "github.com/docker/docker/client"
)

var config = DockerClientConfig{
	Version: "1.39",
}

func Client() (*DockerClient, error) {
	cli, err := ClientFromConfig(config)
	return cli, err
}

func ClientFromConfig(config DockerClientConfig) (*DockerClient, error) {
	cli, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithVersion(config.Version))

	if err != nil {
		return nil, err
	}

	return &DockerClient{Daemon: cli}, nil
}
