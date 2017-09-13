package api

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/swarm"
)

// SwarmInit init swarm if needed
func SwarmInit(config SwarmConfig) (string, error) {
	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return "", err
	}

	s := &swarm.InitRequest{
		ListenAddr:       fmt.Sprintf("%s:%s", config.PrivateIP, config.SwarmPort),
		AdvertiseAddr:    fmt.Sprintf("%s:%s", config.PrivateIP, config.SwarmPort),
		ForceNewCluster:  false,
		AutoLockManagers: false,
	}

	_, err = client.SwarmInit(ctx, *s)
	if err != nil {
		return "", err
	}
	return "Swarm cluster initialized", nil
}
