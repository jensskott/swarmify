package ovh

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/rackspace/gophercloud"
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

func waitForIP(client *gophercloud.ServiceClient, id string) ([]string, error) {

	var ips []string
	nets := []string{"VLAN-Static", "Ext-Net"}

	for _, i := range nets {
		var ipData IPAddress

		for {
			srv, _ := servers.Get(client, id).Extract()
			if srv.Addresses[i] == nil {
				time.Sleep(time.Second * 20)
				continue
			}
			fmt.Println(srv.Addresses[i])
			jsonByte, err := json.Marshal(srv.Addresses[i])
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(jsonByte, &ipData)
			if err != nil {
				return nil, err
			}

			ips = append(ips, ipData.Addr)

			if len(ips) > 2 {
				break
			}
		}
	}

	return ips, nil
}
