package api

import (
	"github.com/docker/docker/client"
)

// Connect to docker client
func Connect(config SwarmConfig) (*client.Client, error) {
	c, err := client.NewClient(config.Endpoint, "1.30", nil, nil)
	if err != nil {
		return nil, err
	}
	return c, nil
}
