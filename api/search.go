package api

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
)

// SearchSwarmCluster for the correct nodes
func SearchSwarmCluster(config SwarmConfig) ([]string, error) {
	var nodeData []string

	ctx := context.Background()

	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	f := filters.NewArgs()
	f.Add("role", config.Nodetype)

	l := &types.NodeListOptions{
		Filters: f,
	}

	nodes, _ := client.NodeList(ctx, *l)
	if nodes != nil {
		nodeData = append(nodeData, nodes[0].ManagerStatus.Addr)
		nodeData = append(nodeData, fmt.Sprintf("%s", nodes[0].Status.State))
	}

	return nodeData, nil
}
