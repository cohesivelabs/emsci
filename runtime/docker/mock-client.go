package docker

import (
	"bytes"
	"context"
	"errors"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	uuid "github.com/gofrs/uuid"
	"io"
	"io/ioutil"
)

type MockDockerClient struct {
	images              map[string]string
	createdContainers   map[string]string
	startedContainers   map[string]string
	PreExecuteCallbacks map[string]func(...interface{})
}

func NewMockDockerClient() *MockDockerClient {
	image1Id, _ := uuid.NewV4()
	image2Id, _ := uuid.NewV4()

	images := map[string]string{
		"image-1:latest": image1Id.String(),
		"image-2:1.0":    image2Id.String(),
	}

	return &MockDockerClient{
		images:            images,
		createdContainers: make(map[string]string),
		startedContainers: make(map[string]string),
	}
}

func (client *MockDockerClient) Info(ctx context.Context) (types.Info, error) {
	var err error

	result := types.Info{
		ServerVersion: "1.38",
	}

	if preexec, ok := client.PreExecuteCallbacks["Info"]; ok {
		preexec(client, ctx, &result, &err)
	}

	return result, err
}

func (client *MockDockerClient) ClientVersion() string {
	result := "1.38"

	if preexec, ok := client.PreExecuteCallbacks["ClientVersion"]; ok {
		preexec(client, &result)
	}

	return result
}

func (client *MockDockerClient) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	var err error

	newImages := client.images

	if _, exists := newImages[refStr]; !exists {
		imageId, _ := uuid.NewV4()
		newImages[refStr] = imageId.String()
	}

	if preexec, ok := client.PreExecuteCallbacks["ImagePull"]; ok {
		preexec(client, ctx, refStr, client.images, &newImages, &err)
	}

	client.images = newImages

	return ioutil.NopCloser(bytes.NewBuffer([]byte("this is a test"))), err
}

func (client *MockDockerClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	var err error
	mappedImages := make([]types.ImageSummary, 0)

	for k, v := range client.images {
		mappedImages = append(mappedImages, types.ImageSummary{
			ID:       v,
			RepoTags: []string{k},
		})
	}

	if preexec, ok := client.PreExecuteCallbacks["ImageList"]; ok {
		preexec(client, ctx, client.images, &mappedImages, &err)
	}

	return mappedImages, err
}

func (client *MockDockerClient) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	var err error

	mappedImages := make([]types.ImageDeleteResponseItem, 0)

	for _, v := range client.images {
		if v == imageID {
			mappedImages = append(mappedImages, types.ImageDeleteResponseItem{
				Deleted: v,
			})
		}
	}

	if preexec, ok := client.PreExecuteCallbacks["ImageRemove"]; ok {
		preexec(client, ctx, client.images, &mappedImages, &err)
	}

	return mappedImages, err
}

func (client *MockDockerClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	var err error
	var containerId uuid.UUID

	containerId, err = uuid.NewV4()
	newContainer := &container.ContainerCreateCreatedBody{
		ID: containerId.String(),
	}

	newContainers := client.createdContainers

	newContainers[containerId.String()] = containerName

	if preexec, ok := client.PreExecuteCallbacks["ContainerCreate"]; ok {
		preexec(client, ctx, containerName, client.createdContainers, containerId, containerName, newContainer, &newContainers, &err)
	}

	client.createdContainers = newContainers

	return *newContainer, err
}

func (client *MockDockerClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	var err error

	if _, exists := client.createdContainers[containerID]; !exists {
		return errors.New("container not created")
	}

	newContainers := client.startedContainers

	newContainers[containerID] = client.createdContainers[containerID]

	if preexec, ok := client.PreExecuteCallbacks["ContainerStart"]; ok {
		preexec(client, ctx, client.createdContainers, client.startedContainers, containerID, &newContainers, &err)
	}

	client.startedContainers = newContainers

	return err
}

func (client *MockDockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	var err error
	containers := make([]types.Container, 0)

	for k, v := range client.startedContainers {
		containers = append(containers, types.Container{
			ID:    k,
			Names: []string{"/" + v},
		})
	}

	if preexec, ok := client.PreExecuteCallbacks["ContainerList"]; ok {
		preexec(client, ctx, &containers, &err)
	}

	return containers, err
}

func (client *MockDockerClient) ContainerKill(ctx context.Context, containerID, signal string) error {
	var err error
	cc := make(map[string]string)
	sc := make(map[string]string)

	for k, v := range client.createdContainers {
		if k != containerID {
			cc[k] = v
		}
	}

	for k, v := range client.startedContainers {
		if k != containerID {
			sc[k] = v
		}
	}

	if preexec, ok := client.PreExecuteCallbacks["ContainerKill"]; ok {
		preexec(client, ctx, containerID, &cc, &sc, &err)
	}

	return err
}
