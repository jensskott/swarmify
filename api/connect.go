package api

import (
	docker "github.com/fsouza/go-dockerclient"
)

// Connect to docker client
func Connect(config SwarmConfig) (*docker.Client, error) {

	c, err := docker.NewClient(config.Endpoint)
	if err != nil {
		return nil, err
	}
	return c, nil
}
