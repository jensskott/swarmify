package api

import (
	"context"
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

// SearchSwarmCluster for the correct nodes
func SearchSwarmCluster(config SwarmConfig) ([]string, error) {
	var nodeData []string

	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	x := make(map[string][]string)

	x["role"] = append(x["role"], config.Nodetype)

	l := &docker.ListNodesOptions{
		Filters: x,
		Context: ctx,
	}

	nodes, _ := client.ListNodes(*l)
	if nodes != nil {
		nodeData = append(nodeData, nodes[0].ManagerStatus.Addr)
		nodeData = append(nodeData, fmt.Sprintf("%s", nodes[0].Status.State))
	}

	return nodeData, nil
}
