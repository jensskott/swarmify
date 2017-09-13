package ovh

import (
	"encoding/json"
	"strings"

	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
)

// SearchSwarm by id
func SearchSwarm(config Config, nodetype string) ([]string, error) {

	var ip []string
	var ipData []PrivateIPAddress

	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	opts := servers.ListOpts{Image: "6c3a0a48-981a-4ccc-8d5e-e00c4dc3c4aa"}
	pager := servers.List(client, opts)

	// Define an anonymous function to be executed on each page's iteration
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, _ := servers.ExtractServers(page)
		for _, s := range serverList {
			if strings.Contains(s.Name, nodetype) {
				jsonByte, err := json.Marshal(s.Addresses["VLAN-Static"])
				if err != nil {
					return false, err
				}

				err = json.Unmarshal(jsonByte, &ipData)
				if err != nil {
					return false, err
				}

				ip = append(ip, ipData[0].Addr)
			}
		}
		return true, nil
	})
	if err != nil {
		return nil, err
	}
	return ip, nil
}
