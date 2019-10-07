package runtime

import (
	"context"
)

type DockerClient struct{}

type ContainerFeatures interface {
	ContainerCreate(ctx context.Context, imageName string, containerName string) (id string, err error)
	ContainerGetByID(ctx context.Context, id string) (*Container, error)
	ContainerGetByName(ctx context.Context, name string) (*Container, error)
	ContainerList(ctx context.Context) ([]Container, error)
	ContainerRemove(ctx context.Context, id string) error
}

type ImageFeatures interface {
	ImagePull(ctx context.Context, imagName string) error
	ImageFind(ctx context.Context, imageName string) (id *string, err error)
	ImageExists(ctx context.Context, imageName string) (bool, error)
	ImageDelete(ctx context.Context, imageName string) error
}

type Runtimer interface {
	ContainerFeatures
	ImageFeatures
	ServerVersion(ctx context.Context) (string, error)
	ClientVersion() string
}

type Port struct {
	IP          string
	PrivatePort uint16
	PublicPort  uint16
	Type        string
}

type Container struct {
	ID      string
	Names   []string
	Image   string
	ImageID string
	Created int64
	Ports   []Port
	State   string
	Status  string
}
