package ovh

import (
	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
)

func CreateCompute(config Config, nodetype string) ([]string, error) {
	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	network1 := servers.Network{UUID: config.Networks[0]}
	network2 := servers.Network{UUID: config.Networks[1]}

	networks := []servers.Network{network1, network2}

	name := ("swarm" + nodetype)

	serverOpts := &servers.CreateOpts{
		Name:       name,
		FlavorName: config.FlavorName,
		ImageRef:   config.ImageID,
		Networks:   networks,
	}

	server, err := servers.Create(client, *serverOpts).Extract()
	if err != nil {
		return nil, err
	}

	ips, err := waitForIP(client, server.ID)
	if err != nil {
		return nil, err
	}

	// Just when testing
	servers.Delete(client, server.ID)

	return ips, nil
}
