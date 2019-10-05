package docker

import (
	"context"
	dockerApi "github.com/docker/docker/client"
)

var config = ClientConfig{
	Version: "1.39",
}

var defaultClient = DockerClient{}

func NewClient() (DockerClient, error) {
	cli, err := NewClientFromConfig(config)
	return cli, err
}

func NewClientFromConfig(config ClientConfig) (DockerClient, error) {
	api, err := dockerApi.NewClientWithOpts(dockerApi.FromEnv, dockerApi.WithVersion(config.Version))

	if err != nil {
		return defaultClient, err
	}

	return DockerClient{Api: api}, nil
}

func (client DockerClient) ServerVersion (ctx context.Context) (string, error) {
	info, err := client.Api.Info(ctx)
	if err != nil {
		return "", err
	}

	return info.ServerVersion, nil
}

func (client DockerClient) ClientVersion () string {
	return client.Api.ClientVersion()
}
