package tests

import (
	"reflect"
	"testing"
	docker "emsci/runtime/docker"
	runtimeTypes "emsci/runtime/types"
)

var mockDockerClient MockDockerClient = MockDockerClient{}

//TODO: move these to an intergration test suite
func TestNewClient(t *testing.T) {
	t.Run("Should return a docker client", func(t *testing.T) {
		client := docker.DockerClient{
			Api: mockDockerClient,
		}

		_, ok := reflect.TypeOf(client).(runtimeTypes.Runtimer)
		if(!ok) {
			t.Error("client does not implement Runtimer interface")
		}
	})
}
