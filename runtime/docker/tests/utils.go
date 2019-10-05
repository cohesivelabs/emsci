package tests

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"io"
)

type MockDockerClient struct {
	info func(ctx context.Context) (types.Info, error)
	clientVersion func() string
	imagePull func(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)
	imageList func(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)
	imageRemove func(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)
	containerCreate func(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)
	containerStart func(ctx context.Context, containerID string, options types.ContainerStartOptions) error
	containerList func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)
	containerKill func(ctx context.Context, containerID, signal string) error
	images []string
	containers []string
}

func NewMockDockerClient() MockDockerClient {
	info := func(ctx context.Context) (types.Info, error) {
		return types.Info{
			ServerVersion: "1.38",
		}, nil
	}

	clientVersion := func() string {
		return "1.38"
	}

	imagePull := func(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
		return nil, nil
	}

	imageList := func(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
		return []types.ImageSummary{
			types.ImageSummary{
				ID: "image-1",
				RepoTags: []string {
					"latest",
					"1.1",
					"1.2",
				},
			},
			types.ImageSummary{
				ID: "image-2",
				RepoTags: []string {
					"latest",
					"1.0",
					"2.1",
				},
			},
		}, nil
	}

	imageRemove := func(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
		return []types.ImageDeleteReponseItem{
			types.ImageDeleteResponseItem{},
		}, nil
	}

	containerCreate := func(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
		return container.ContainerCreateCreatedBody{
			ID: "container-created-" + containerName,
		}, nil
	}

	containerStart := func(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
		return nil
	}

	containerList := func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	}

	containerKill := func(ctx context.Context, containerID, signal string) error

	return MockDockerClient{
	}
}

// OverrideInfo provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideInfo(override func(ctx context.Context) (types.Info, error)) {
	client.info = override
}

// OverrideClientVersion provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideClientVersion (override func()(string)) {
	client.clientVersion = override
}

// OverrideImagePull provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideImagePull(override func(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error)) {
	client.imagePull = override
}

// OverrideImageList provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideImageList(override func(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error)) {
	client.imageList = override
}

// OverrideImageRemove provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideImageRemove(override func(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error)) {
	client.imageRemove = override
}

// OverrideContainerCreate provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideContainerCreate(override func(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error)) {
	client.containerCreate = override
}

// OverrideContainerStart provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideContainerStart(override func(ctx context.Context, containerID string, options types.ContainerStartOptions) error) {
	client.containerStart = override
}

// OverrideContainerList provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideContainerList(override func(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error)) {
	client.containerList = override
}

// OverrideContainerKill provides an escape hatch to allow the test to control exactly what this method does
func (client *MockDockerClient) OverrideContainerKill(override func(ctx context.Context, containerID, signal string) error) {
	client.containerKill = override
}

func (client *MockDockerClient) Info(ctx context.Context) (types.Info, error) {
	return client.info(ctx)
}

func (client *MockDockerClient) ClientVersion () string {
	return client.clientVersion()
}

func (client *MockDockerClient) ImagePull(ctx context.Context, refStr string, options types.ImagePullOptions) (io.ReadCloser, error) {
	return client.imagePull(ctx, refStr, options)
}

func (client *MockDockerClient) ImageList(ctx context.Context, options types.ImageListOptions) ([]types.ImageSummary, error) {
	return client.imageList(ctx, options)
}

func (client *MockDockerClient) ImageRemove(ctx context.Context, imageID string, options types.ImageRemoveOptions) ([]types.ImageDeleteResponseItem, error) {
	return client.imageRemove(ctx, imageID, options)
}

func (client *MockDockerClient) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	return client.containerCreate(ctx, config, hostConfig, networkingConfig, containerName)
}

func (client *MockDockerClient) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	return client.containerStart(ctx, containerID, options)
}

func (client *MockDockerClient) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return client.containerList(ctx, options)
}

func (client *MockDockerClient) ContainerKill(ctx context.Context, containerID, signal string) error {
	return client.containerKill(ctx, containerID, signal)
}
