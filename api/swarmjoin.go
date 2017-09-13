package api

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/swarm"
	docker "github.com/fsouza/go-dockerclient"
)

// JoinSwarm with token
func JoinSwarm(config SwarmConfig) (string, error) {
	ctx := context.Background()

	var token string

	client, err := Connect(config)
	if err != nil {
		return "", err
	}

	if config.Nodetype == "manager" {
		token = config.Managertoken
	} else {
		token = config.Workertoken
	}

	join := &docker.JoinSwarmOptions{
		swarm.JoinRequest{
			ListenAddr:    config.PrivateIP,
			AdvertiseAddr: config.ClientIP,
			RemoteAddrs:   config.SwarmMaster,
			JoinToken:     token,
		},
		ctx,
	}

	err = client.JoinSwarm(*join)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Node joined to swarm as %s\n", config.Nodetype), nil
}
