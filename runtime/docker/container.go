package docker

import (
	"context"
	dockerTypes "github.com/docker/docker/api/types"
	containerTypes "github.com/docker/docker/api/types/container"
	runtimeTypes "emsci/runtime/types"
	"emsci/log"
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

	container, err := client.Api.ContainerCreate(ctx, containerConfig, nil, nil, containerName)
	if err != nil {
		return "", err
	}

	for _, warning := range container.Warnings {
		log.Warn(warning)
	}

	if err := client.Api.ContainerStart(ctx, container.ID, dockerTypes.ContainerStartOptions{}); err != nil {
		return "", err
	}

	return container.ID, nil
}

// ContainerGetByID - return a container instance for a given id
func (client DockerClient) ContainerGetByID(ctx context.Context, id string) (*runtimeTypes.Container, error) {
	var container *runtimeTypes.Container = nil
	var err error

	if containers, err := client.ContainerList(ctx); err == nil {
		for _, item := range containers {
			if item.ID == id {
				container = &item
				break
			}
		}
	}

	return container, err
}

// ContainerGetByName - return a container instance for a given name
func (client DockerClient) ContainerGetByName(ctx context.Context, name string) (*runtimeTypes.Container, error) {
	var container *runtimeTypes.Container = nil
	var err error

	if containers, err := client.ContainerList(ctx); err == nil {
		for _, item := range containers {
			for _, x := range item.Names {
				if x == "/"+name {
					container = &item
					break
				}
			}
		}
	}

	return container, err
}

// ContainerList - returns a list of all running containers
func (client DockerClient) ContainerList(ctx context.Context) ([]runtimeTypes.Container, error) {
	runtimeContainers := make([]runtimeTypes.Container, 0)
	var err error

	if containers, err := client.Api.ContainerList(ctx, dockerTypes.ContainerListOptions{}); err == nil {
		for _, dc := range containers {
			ports := make([]runtimeTypes.Port, 0)

			for _, p := range dc.Ports {
				ports = append(ports, runtimeTypes.Port{
					IP: p.IP,
					PrivatePort: p.PrivatePort,
					PublicPort: p.PublicPort,
					Type: p.Type,
				})
			}

			runtimeContainers = append(runtimeContainers, runtimeTypes.Container{
				ID: dc.ID,
				Names: dc.Names,
				Image: dc.Image,
				ImageID: dc.ImageID,
				Created: dc.Created,
				State: dc.State,
				Status: dc.Status,
				Ports: ports,
			})
		}
	}

	return runtimeContainers, err
}

func (client DockerClient) ContainerRemove(ctx context.Context, id string) error {
	if err := client.Api.ContainerKill(ctx, id, "SIGKILL"); err != nil {
		return err
	}

	return nil
}
