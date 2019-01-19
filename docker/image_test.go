package docker

import (
	"context"
	"testing"
)

func TestDockerClient_ImagePull(t *testing.T) {
	client, err := Client()
	if err != nil {
		t.Error(err)
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

func BenchmarkDockerClient_ImagePull(t *testing.B) {
	client, err := Client()
	if err != nil {
		t.Error(err)
	}

	imageName := "busybox:latest"
	ctx := context.Background()

	defer client.ImageDelete(ctx, imageName)

	client.ImagePull(ctx, imageName)

	_, err = client.ImageExists(ctx, imageName)
	if err != nil {
		t.Error(err)
	}
}
