package ovh

import (
	"encoding/json"
	"strings"

	"github.com/rackspace/gophercloud/openstack/compute/v2/servers"
	"github.com/rackspace/gophercloud/pagination"
)

// PrivateIPAddress data
type PrivateIPAddress struct {
	OSEXTIPSMACMacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	OSEXTIPSType       string `json:"OS-EXT-IPS:type"`
	Addr               string `json:"addr"`
	Version            int    `json:"version"`
}

// SearchSwarm by id
func SearchSwarm(config Config, nodetype string) ([]string, error) {

	var ip []string
	var ipData []PrivateIPAddress

	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	opts := servers.ListOpts{Image: "a4564ff3-d226-422e-98fe-b0c753dd4657"}
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
