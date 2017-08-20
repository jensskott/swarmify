package api

import "context"

// SwarmTokens from manager
func SwarmTokens(config SwarmConfig) (map[string]string, error) {

	var tokens map[string]string
	tokens = make(map[string]string)

	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	x, _ := client.InspectSwarm(ctx)
	tokens["Manager"] = x.JoinTokens.Manager
	tokens["Worker"] = x.JoinTokens.Worker

	return tokens, nil
}
