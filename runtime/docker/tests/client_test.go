package tests

import (
	docker "emsci/runtime/docker"
	runtimeTypes "emsci/runtime/types"
	"testing"
)

//TODO: move these to an intergration test suite
func TestNewClient(t *testing.T) {
	t.Run("Should return a docker client", func(t *testing.T) {
		client := docker.DockerClient{
			Api: NewMockDockerClient(),
		}

		_, ok := interface{}(client).(runtimeTypes.Runtimer)
		if !ok {
			t.Error("client does not implement Runtimer interface")
		}
	})
}
