package api

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
)

func RunSwarmConfig(nodetype, endpoint string) (string, error) {
	ctx := context.Background()

	ep := endpoint
	client, err := docker.NewClient(ep)
	if err != nil {
		return "", err
	}

	if nodetype == "manager" {
		node := SearchSwarmCluster(ctx, client, nodetype)
		if node == nil {
			err = SwarmInit(ctx, client)
			if err != nil {
				return "", err
			}
			return "Docker swarm cluster initiated", nil
		}
		tkn := SwarmManagerToken(ctx, client)
		return tkn, nil
	}

	return "", nil
}
