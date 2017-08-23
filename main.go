package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	api "github.com/jensskott/swarmify/api"
	"github.com/jensskott/swarmify/ovh"
)

// DockerConfigFile for the run
type DockerConfigFile struct {
	Endpoint     string `yaml:"endpoint"`
	Nodetype     string `yaml:"nodetype"`
	SwarmPort    string `yaml:"swarmport"`
	WorkerToken  string `yaml:"workertoken"`
	ManagerToken string `yaml:"managertoken"`
}

// OvhConfigFile for the run
type OvhConfigFile struct {
	IdentityEndpoint string `yaml:"identityendpoint"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	TenantID         string `yaml:"tenantid"`
	DomainName       string `yaml:"domainname"`
	Region           string `yaml:"region"`
}

// AppConfig for the whole app
type AppConfig struct {
	DockerConfig DockerConfigFile
	OvhConfig    OvhConfigFile
	PrivateIP    string
	ClientIP     string
}

func main() {

	var x DockerConfigFile
	var y OvhConfigFile

	dir, _ := os.Getwd()
	dockerYaml := (dir + "/docker.yaml")
	ovhYaml := (dir + "/ovh.yaml")

	// Read config from yaml file
	dockerYamlFile, err := ioutil.ReadFile(dockerYaml)
	check(err)

	ovhYamlFile, err := ioutil.ReadFile(ovhYaml)
	check(err)

	err = yaml.Unmarshal(dockerYamlFile, &x)
	check(err)

	err = yaml.Unmarshal(ovhYamlFile, &y)

	config := &AppConfig{
		DockerConfig: x,
		OvhConfig:    y,
		PrivateIP:    os.Args[2],
		ClientIP:     os.Args[2],
	}

	ovhCfg := &ovh.Config{
		IdentityEndpoint: config.OvhConfig.IdentityEndpoint,
		Username:         config.OvhConfig.Username,
		Password:         config.OvhConfig.Password,
		TenantID:         config.OvhConfig.TenantID,
		DomainName:       config.OvhConfig.DomainName,
		Region:           config.OvhConfig.Region,
	}

	masterIPs, err := ovh.SearchSwarm(*ovhCfg, config.DockerConfig.Nodetype)
	check(err)

	dockerCfg := &api.SwarmConfig{
		Endpoint:     config.DockerConfig.Endpoint,
		Nodetype:     config.DockerConfig.Nodetype,
		SwarmPort:    config.DockerConfig.SwarmPort,
		SwarmMaster:  masterIPs,
		Managertoken: config.DockerConfig.ManagerToken,
		Workertoken:  config.DockerConfig.WorkerToken,
		PrivateIP:    config.PrivateIP,
		ClientIP:     config.ClientIP,
	}

	switch os.Args[1] {
	case "init":
		resp, err := api.SwarmInit(*dockerCfg)
		check(err)

		token, err := api.SwarmTokens(*dockerCfg)
		check(err)

		z := &DockerConfigFile{
			Endpoint:     config.DockerConfig.Endpoint,
			Nodetype:     config.DockerConfig.Nodetype,
			SwarmPort:    config.DockerConfig.SwarmPort,
			ManagerToken: token["Manager"],
			WorkerToken:  token["Worker"],
		}

		yaml, err := yaml.Marshal(*z)
		check(err)

		err = ioutil.WriteFile(dockerYaml, yaml, 0644)
		check(err)

		log.Println(resp)
	case "manager", "worker":

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
