package api

import (
	"context"

	"github.com/docker/docker/api/types/swarm"
	docker "github.com/fsouza/go-dockerclient"
)

// SwarmInit init swarm if needed
func SwarmInit(config SwarmConfig) error {
	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return err
	}

	s := &docker.InitSwarmOptions{
		swarm.InitRequest{
			ListenAddr:       "0.0.0.0:2377",
			AdvertiseAddr:    "127.0.0.1:2377",
			ForceNewCluster:  false,
			AutoLockManagers: false,
		},
		ctx,
	}
	_, err = client.InitSwarm(*s)
	if err != nil {
		return err
	}
	return nil
}
