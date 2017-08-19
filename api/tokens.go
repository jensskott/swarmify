package api

import "context"

// SwarmManagerToken lookup
func SwarmManagerToken(config SwarmConfig) (string, error) {
	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return "", err
	}

	x, _ := client.InspectSwarm(ctx)
	return x.JoinTokens.Manager, nil
}

// SwarmWorkerToken lookup
func SwarmWorkerToken(config SwarmConfig) (string, error) {
	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return "", err
	}

	x, _ := client.InspectSwarm(ctx)
	return x.JoinTokens.Worker, nil
}
