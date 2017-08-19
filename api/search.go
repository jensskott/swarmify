package api

import (
	"context"
	"fmt"

	docker "github.com/fsouza/go-dockerclient"
)

func SearchSwarmCluster(ctx context.Context, client *docker.Client, role string) []string {
	var nodeData []string

	x := make(map[string][]string)

	x["role"] = append(x["role"], role)

	l := &docker.ListNodesOptions{
		Filters: x,
		Context: ctx,
	}

	nodes, _ := client.ListNodes(*l)
	if nodes != nil {
		nodeData = append(nodeData, nodes[0].ManagerStatus.Addr)
		nodeData = append(nodeData, fmt.Sprintf("%s", nodes[0].Status.State))
	}

	return nodeData
}
