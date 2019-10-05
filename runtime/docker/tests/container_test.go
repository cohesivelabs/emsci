package tests

import (
	"context"
	"testing"
	"emsci/runtime/docker"

	uuid "github.com/gofrs/uuid"
)

func TestDockerClient_ContainerCreate(t *testing.T) {
	t.Run("should be able to create a container", func(t *testing.T) {

		client := docker.DockerClient{
			Api: MockDockerClient{},
		}

		imageName := "busybox:latest"
		ctx := context.Background()

		containerName := "container-create-test-1-" + uuid.Must(uuid.NewV4()).String()
		result, err := client.ContainerCreate(ctx, imageName, containerName)
		if err != nil {
			t.Error(err)
		}

		defer client.ContainerRemove(ctx, containerName)

		container, err := client.ContainerGetByID(ctx, result)
		if err != nil {
			t.Error(err)
		}

		if container == nil {
			t.Errorf("Container with ID %v was not found", result)
		}
	})
}

func TestDockerClient_ContainerGetByName(t *testing.T) {
	t.Run("Should be able to get container by name", func(t *testing.T) {
		containerName := "container-getByName-test-1-" + uuid.Must(uuid.NewV4()).String()

		client := docker.DockerClient{
			Api: mockDockerClient,
		}

		imageName := "busybox:latest"

		ctx := context.Background()
		defer client.ImageDelete(ctx, imageName)

		if err := client.ImagePull(ctx, imageName); err != nil {
			t.Error(err)
		}

		result, err := client.ContainerCreate(ctx, imageName, containerName)
		if err != nil {
			t.Error(err)
		}

		defer client.ContainerRemove(ctx, result)

		container, err := client.ContainerGetByName(ctx, containerName)
		if err != nil {
			t.Error(err)
		}

		if container == nil {
			t.Errorf("Container with name %v was not found", containerName)
		}
	})
}

func TestDockerClient_ContainerList(t *testing.T) {
	t.Run("should be able to get a list of containers", func(t *testing.T) {
		client := docker.DockerClient{
			Api: mockDockerClient,
		}

		imageName := "busybox:latest"
		ctx := context.Background()

		client.ImagePull(ctx, imageName)
		defer client.ImageDelete(ctx, imageName)

		container1Ch := make(chan string)
		container2Ch := make(chan string)
		errorCh := make(chan error)

		containerName1 := "container-list-test-1-" + uuid.Must(uuid.NewV4()).String()
		containerName2 := "container-list-test-2-" + uuid.Must(uuid.NewV4()).String()

		go func() {
			container, err := client.ContainerCreate(ctx, imageName, containerName1)
			if err != nil {
				errorCh <- err
				return
			}

			container1Ch <- container
		}()

		go func() {
			container, err := client.ContainerCreate(ctx, imageName, containerName2)
			if err != nil {
				errorCh <- err
				return
			}

			container2Ch <- container
		}()

		var container1 string
		var container2 string

		select {
		case err := <-errorCh:
			t.Error(err)
		case container := <-container1Ch:
			container1 = container
		}

		select {
		case err := <-errorCh:
			t.Error(err)
		case container := <-container2Ch:
			container2 = container
		}

		defer client.ContainerRemove(ctx, container1)
		defer client.ContainerRemove(ctx, container2)

		containers, err := client.ContainerList(ctx)
		if err != nil {
			t.Error(err)
		}

		if len(containers) == 0 {
			t.Error("Container list empty")
		}
	})
}
