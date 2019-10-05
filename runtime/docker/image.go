package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"io/ioutil"
)

func (client DockerClient) ImagePull(ctx context.Context, imageName string) error {
	options := types.ImagePullOptions{}

	resp, err := client.Api.ImagePull(ctx, imageName, options)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(resp)
	if err != nil {
		return err
	}

	return nil
}

func (client DockerClient) ImageFind(ctx context.Context, imageName string) (id *string, err error) {
	images, err := client.Api.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	for _, image := range images {
		if id != nil {
			break
		}

		for _, tag := range image.RepoTags {
			if tag == imageName {
				id = &image.ID
				break
			}
		}
	}

	return id, nil
}

func (client DockerClient) ImageExists(ctx context.Context, imageName string) (bool, error) {
	id, err := client.ImageFind(ctx, imageName)

	if err != nil {
		return false, err
	}

	if id != nil {
		return true, nil
	}

	return false, nil
}

func (client DockerClient) ImageDelete(ctx context.Context, imageName string) error {
	id, err := client.ImageFind(ctx, imageName)

	if err != nil || id == nil {
		fmt.Printf("Could not delete image %s because it was not found or an error occured", imageName)
		fmt.Print(err)
	}

	_, err = client.Api.ImageRemove(ctx, imageName, types.ImageRemoveOptions{Force: true, PruneChildren: true})
	if err != nil {
		return err
	}

	return nil
}
