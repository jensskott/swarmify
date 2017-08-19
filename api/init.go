package api

import (
	"context"

	"github.com/docker/docker/api/types/swarm"
	docker "github.com/fsouza/go-dockerclient"
)

func SwarmInit(ctx context.Context, client *docker.Client) error {
	s := &docker.InitSwarmOptions{
		swarm.InitRequest{
			ListenAddr:       "0.0.0.0:2377",
			AdvertiseAddr:    "127.0.0.1:2377",
			ForceNewCluster:  false,
			AutoLockManagers: false,
		},
		ctx,
	}
	_, err := client.InitSwarm(*s)
	if err != nil {
		return err
	}
	return nil
}
