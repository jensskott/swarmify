package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/jensskott/swarmify/api"
	"github.com/jensskott/swarmify/ovh"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/networking/v2/networks"
	"github.com/rackspace/gophercloud/pagination"
)

func main() {

	dir, _ := os.Getwd()
	dockerYaml := (dir + "/docker.yaml")
	ovhYaml := (dir + "/ovh.yaml")

	dockerStruct, ovhStruct := readYaml(dockerYaml, ovhYaml)

	config := &AppConfig{
		DockerConfig: dockerStruct,
		OvhConfig:    ovhStruct,
	}

	switch os.Args[1] {
	case "init":
		ovhCfg := &ovh.Config{
			IdentityEndpoint: config.OvhConfig.IdentityEndpoint,
			Username:         config.OvhConfig.Username,
			Password:         config.OvhConfig.Password,
			TenantID:         config.OvhConfig.TenantID,
			DomainName:       config.OvhConfig.DomainName,
			Region:           config.OvhConfig.Region,
			ImageID:          config.OvhConfig.ImageID,
			FlavorName:       config.OvhConfig.FlavorName,
			Count:            1,
			Networks:         config.OvhConfig.Networks,
		}

		computeResp, err := ovh.CreateCompute(*ovhCfg, "manager")
		check(err)

		fmt.Println(computeResp)

		dockerCfg := &api.SwarmConfig{
			Endpoint:  config.DockerConfig.Endpoint,
			SwarmPort: config.DockerConfig.SwarmPort,
			PrivateIP: "127.0.0.1",
			ClientIP:  "127.0.0.1",
		}

		initResp, err := api.SwarmInit(*dockerCfg)
		check(err)

		log.Println(initResp)

		token, err := api.SwarmTokens(*dockerCfg)
		check(err)

		z := &DockerConfigFile{
			Endpoint:     config.DockerConfig.Endpoint,
			Nodetype:     config.DockerConfig.Nodetype,
			SwarmPort:    config.DockerConfig.SwarmPort,
			ManagerToken: token["Manager"],
			WorkerToken:  token["Worker"],
		}

		writeYaml(z, dockerYaml)

	case "manager", "worker":

		/*
			        masterIPs, err := ovh.SearchSwarm(*ovhCfg, config.DockerConfig.Nodetype)
					check(err)

					createResp := ovh.CreateCompute(*ovhCfg, config.DockerConfig.Nodetype)
					resp, err := api.JoinSwarm(*dockerCfg)
					check(err)

			        log.Println(resp)
		*/
	case "heal":
		log.Println("Healing not active anymore")
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func findnetworks() {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: "https://auth.cloud.ovh.net/v3",
		Username:         "A6G9Dta96qxD",
		Password:         "9djNmvwwqF8CaMU9uwn9aRsC7U3YAyMa",
		TenantID:         "6f4784ed5ce5486084ed1004c53c2642",
		DomainName:       "default",
	}
	provider, err := openstack.AuthenticatedClient(opts)
	check(err)

	client, err := openstack.NewNetworkV2(provider, gophercloud.EndpointOpts{
		Region: "BHS3",
	})

	myBool := false

	netOpts := &networks.ListOpts{Shared: &myBool}

	// Retrieve a pager (i.e. a paginated collection)
	pager := networks.List(client, netOpts)

	// Define an anonymous function to be executed on each page's iteration
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		networkList, _ := networks.ExtractNetworks(page)

		for _, n := range networkList {
			fmt.Printf("Network: %v  ID: %v", n.Name, n.ID)
		}
		return false, nil
	})
}
