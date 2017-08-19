package api

import (
	"context"

	docker "github.com/fsouza/go-dockerclient"
)

func SwarmManagerToken(ctx context.Context, client *docker.Client) string {
	x, _ := client.InspectSwarm(ctx)
	return x.JoinTokens.Manager
}

func SwarmWorkerToken(ctx context.Context, client *docker.Client) string {
	x, _ := client.InspectSwarm(ctx)
	return x.JoinTokens.Worker
}
