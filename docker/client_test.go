package docker

import (
	"reflect"
	"testing"

	docker "github.com/docker/docker/client"
)

func TestClient(t *testing.T) {
	t.Run("Should return a docker client", func(t *testing.T) {
		client, err := Client()
		if err != nil {
			t.Error(err)
		}

		switch reflect.TypeOf(client.Daemon) {
		case reflect.TypeOf(&docker.Client{}):
			break
		default:
			t.Error("did not return a docker client")
		}
	})
}

func TestClientFromConfig(t *testing.T) {
	t.Run("Should return a docker client from config", func(t *testing.T) {
		expectedVersion := "1.38"
		config := DockerClientConfig{Version: expectedVersion}

		client, err := ClientFromConfig(config)
		if err != nil {
			t.Error(err)
		}

		switch reflect.TypeOf(client.Daemon) {
		case reflect.TypeOf(&docker.Client{}):
			break
		default:
			t.Error("did not return a docker client")
		}

		actualVersion := client.Daemon.ClientVersion()

		if actualVersion != expectedVersion {
			t.Errorf("did not return correct docker client version: expected %v got %v", expectedVersion, actualVersion)
		}
	})
}
