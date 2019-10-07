package tests

import (
	"context"
	"emsci/runtime/docker"
	"testing"
)

func TestDockerClient_ImagePull(t *testing.T) {
	client := docker.DockerClient{
		Api: NewMockDockerClient(),
	}

	imageName := "busybox:latest"
	ctx := context.Background()

	client.ImagePull(ctx, imageName)

	imageExists, err := client.ImageExists(ctx, imageName)
	if err != nil {
		t.Error(err)
	}

	if imageExists == false {
		t.Fail()
	}
}
