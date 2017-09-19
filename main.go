package main

import (
	cli "gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"

	"fmt"

	"time"

	"github.com/aws/aws-sdk-go/service/clouddirectory"
	api "github.com/jensskott/swarmify/api"
	"github.com/jensskott/swarmify/ovh"
)

var (
	swarmtype  = cli.Flag("type", "What node type to run").Required().String()
	dockerYaml = cli.Flag("dockerconfig", "Docker config file").Required().String()
	ovhYaml    = cli.Flag("ovhconfig", "Ovh config file").Required().String()
)

func main() {
	cli.Version("0.1.0")
	cli.Parse()

	dockerStruct, ovhStruct := readYaml(*dockerYaml, *ovhYaml)

	config := &AppConfig{
		DockerConfig: dockerStruct,
		OvhConfig:    ovhStruct,
	}

	ovhCfg := &ovh.Config{
		IdentityEndpoint: config.OvhConfig.Identityendpoint,
		Username:         config.OvhConfig.Username,
		Password:         config.OvhConfig.Password,
		TenantID:         config.OvhConfig.Tenantid,
		TenantName:       config.OvhConfig.Tenantname,
		DomainName:       config.OvhConfig.Domainname,
		Region:           config.OvhConfig.Region,
		ImageID:          config.OvhConfig.Imageid,
		FlavorName:       config.OvhConfig.Flavorname,
		Rules:            config.OvhConfig.Rules,
	}

	switch *swarmtype {
	case "init":
		infraResp, err := ovh.CreateInfra(*ovhCfg)
		check(err)
		os.Exit(1)

		log.Println(infraResp)

		computeResp, err := ovh.CreateCompute(*swarmtype)
		check(err)

		ep := fmt.Sprintf("http://%s:2376", computeResp[" Ext-Net"])

		// Build docker config for swarm
		dockerCfg := &api.SwarmConfig{
			Endpoint:  ep,
			SwarmPort: config.DockerConfig.SwarmPort,
			PrivateIP: computeResp[" Ext-Net"],
			ClientIP:  computeResp[" Ext-Net"],
		}

		time.Sleep(120 * time.Second)

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
		writeYaml(z, *dockerYaml)

	case "manager", "worker":

		// Search cluster for master ips
		masterIPs, err := ovh.SearchSwarm(*ovhCfg, config.DockerConfig.Nodetype)
		if masterIPs == nil {
			log.Fatal("No ip addresses found")
		}
		check(err)

		// Create compute
		computeResp, err := ovh.CreateCompute(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}

		ep := fmt.Sprintf("http://%s:2376", computeResp[" Ext-Net"])

		// Build docker config for swarm
		dockerCfg := &api.SwarmConfig{
			SwarmMaster:  masterIPs,
			Endpoint:     ep,
			SwarmPort:    config.DockerConfig.SwarmPort,
			PrivateIP:    computeResp[" Ext-Net"],
			ClientIP:     computeResp[" Ext-Net"],
			Managertoken: config.DockerConfig.ManagerToken,
			Workertoken:  config.DockerConfig.WorkerToken,
			Nodetype:     os.Args[1],
		}

		time.Sleep(120 * time.Second)

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
