package api

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/swarm"
	docker "github.com/fsouza/go-dockerclient"
)

// SwarmInit init swarm if needed
func SwarmInit(config SwarmConfig) (string, error) {
	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return "", err
	}

	s := &docker.InitSwarmOptions{
		swarm.InitRequest{
			ListenAddr:       fmt.Sprintf("%s:%s", config.PrivateIP, config.SwarmPort),
			AdvertiseAddr:    fmt.Sprintf("%s:%s", config.PrivateIP, config.SwarmPort),
			ForceNewCluster:  false,
			AutoLockManagers: false,
		},
		ctx,
	}
	_, err = client.InitSwarm(*s)
	if err != nil {
		return "", err
	}
	return "Swarm cluster initialized", nil
}
