package integration

import (
	"context"
	docker "emsci/runtime/docker"
	"testing"
	"time"
)

//TODO: move these to an intergration test suite
func TestNewDockerClient(t *testing.T) {
	t.Run("Should return a docker client", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		client, err := docker.NewClient()
		if err != nil {
			t.Error(err)
		}

		if version, err := client.ServerVersion(ctx); err != nil || version == "" {
			t.Error("client connection was not created successfully")
		}
	})
}

func TestNewDockerClientFromConfig(t *testing.T) {
	t.Run("Should return a docker client from config", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
		defer cancel()

		expectedVersion := "1.38"
		config := docker.ClientConfig{Version: expectedVersion}

		client, err := docker.NewClientFromConfig(config)
		if err != nil {
			t.Error(err)
		}

		actualVersion := client.ClientVersion()

		if actualVersion != expectedVersion {
			t.Errorf("did not return correct docker client version: expected %v got %v", expectedVersion, actualVersion)
		}

		if version, err := client.ServerVersion(ctx); err != nil || version == "" {
			t.Error("client connection was not created successfully")
		}
	})
}
