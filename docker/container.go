package docker

import (
	"context"
	dockerTypes "github.com/docker/docker/api/types"
	containerTypes "github.com/docker/docker/api/types/container"
	"github.com/jmartin84/emsci/log"
)

// ContainerCreate - create a container given an image name
func (client DockerClient) ContainerCreate(ctx context.Context, imageName, containerName string) (string, error) {
	containerConfig := &containerTypes.Config{
		Image:        imageName,
		Tty:          true,
		AttachStderr: false,
		AttachStdin:  false,
		AttachStdout: false,
		StdinOnce:    false,
		OpenStdin:    true,
	}

	if exists, err := client.ImageExists(ctx, imageName); !exists || err != nil {
		if err != nil {
			return "", err
		}

		client.ImagePull(ctx, imageName)
	}

	container, err := client.Daemon.ContainerCreate(ctx, containerConfig, nil, nil, containerName)
	if err != nil {
		return "", err
	}

	for _, warning := range container.Warnings {
		log.Warn(warning)
	}

	if err := client.Daemon.ContainerStart(ctx, container.ID, dockerTypes.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return container.ID, nil
}

// ContainerGetByID - return a container instance for a given id
func (client DockerClient) ContainerGetByID(ctx context.Context, id string) (*dockerTypes.Container, error) {
	var container *dockerTypes.Container = nil

	containers, err := client.ContainerList(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range containers {
		if item.ID == id {
			container = &item
			break
		}
	}

	return container, nil
}

// ContainerGetByName - return a container instance for a given name
func (client DockerClient) ContainerGetByName(ctx context.Context, name string) (*dockerTypes.Container, error) {
	var container *dockerTypes.Container = nil

	containers, err := client.ContainerList(ctx)
	if err != nil {
		return nil, err
	}

	for _, item := range containers {
		for _, x := range item.Names {
			if x == "/"+name {
				container = &item
				break
			}
		}

		if container != nil {
			break
		}
	}

	return container, err
}

// ContainerList - returns a list of all running containers
func (client DockerClient) ContainerList(ctx context.Context) ([]dockerTypes.Container, error) {
	containers, err := client.Daemon.ContainerList(ctx, dockerTypes.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	return containers, nil
}

func (client DockerClient) ContainerRemove(ctx context.Context, id string) error {
	if err := client.Daemon.ContainerKill(ctx, id, "SIGKILL"); err != nil {
		return err
	}

	return nil
}
