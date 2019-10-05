package tests

import (
	"context"
	"testing"
)

var mockDockerClient MockDockerClient = MockDockerClient{}

func TestDockerClient_ImagePull(t *testing.T) {
	client := DockerClient{
		Api: mockDockerClient,
	}


	imageName := "busybox:latest"
	ctx := context.Background()

	defer client.ImageDelete(ctx, imageName)

	client.ImagePull(ctx, imageName)

	imageExists, err := client.ImageExists(ctx, imageName)
	if err != nil {
		t.Error(err)
	}

	if imageExists == false {
		t.Fail()
	}
}
