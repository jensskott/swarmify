package main

import (
	"log"
	"os"

	api "github.com/jensskott/swarmify/api"
	"github.com/jensskott/swarmify/ovh"
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
		computeResp, err := ovh.CreateCompute(config.DockerConfig.Nodetype)
		check(err)

		// Build docker config for swarm
		dockerCfg := &api.SwarmConfig{
			Endpoint:  computeResp[" VLAN-Static"],
			SwarmPort: config.DockerConfig.SwarmPort,
			PrivateIP: computeResp[" VLAN-Static"],
			ClientIP:  computeResp[" VLAN-Static"],
		}

		initResp, err := api.SwarmInit(*dockerCfg)
		check(err)

		// Log out swarm output
		log.Println(initResp)

		// Get the tokens for swarm join
		token, err := api.SwarmTokens(*dockerCfg)
		check(err)

		z := &DockerConfigFile{
			Nodetype:     config.DockerConfig.Nodetype,
			SwarmPort:    config.DockerConfig.SwarmPort,
			ManagerToken: token["Manager"],
			WorkerToken:  token["Worker"],
		}

		// Write new dockerfile with tokens
		writeYaml(z, dockerYaml)

	case "manager", "worker":
		ovhCfg := &ovh.Config{
			IdentityEndpoint: config.OvhConfig.IdentityEndpoint,
			Username:         config.OvhConfig.Username,
			Password:         config.OvhConfig.Password,
			TenantID:         config.OvhConfig.TenantID,
			TenantName:       config.OvhConfig.TenantName,
			DomainName:       config.OvhConfig.DomainName,
			Region:           config.OvhConfig.Region,
			ImageID:          config.OvhConfig.ImageID,
			FlavorName:       config.OvhConfig.FlavorName,
			Count:            "1",
		}

		// Search cluster for master ips
		masterIPs, err := ovh.SearchSwarm(*ovhCfg, config.DockerConfig.Nodetype)
		check(err)

		// Create compute
		computeResp, err := ovh.CreateCompute(config.DockerConfig.Nodetype)
		if err != nil {
			log.Fatal(err)
		}

		// Build docker config for swarm
		dockerCfg := &api.SwarmConfig{
			SwarmMaster: masterIPs,
			Endpoint:    computeResp[" VLAN-Static"],
			SwarmPort:   config.DockerConfig.SwarmPort,
			PrivateIP:   computeResp[" VLAN-Static"],
			ClientIP:    computeResp[" VLAN-Static"],
		}

		// Join swarm
		resp, err := api.JoinSwarm(*dockerCfg)
		check(err)

		log.Println(resp)

	case "heal":
		log.Println("Healing not active anymore")
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
