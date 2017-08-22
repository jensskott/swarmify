package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"

	api "github.com/jensskott/swarmify/api"
)

// ConfigFile for the run
type ConfigFile struct {
	Endpoint     string `yaml:"endpoint"`
	Nodetype     string `yaml:"nodetype"`
	SwarmPort    string `yaml:"swarmport"`
	WorkerToken  string `yaml:"workertoken"`
	ManagerToken string `yaml:"managertoken"`
}

// AppConfig for the whole app
type AppConfig struct {
	DockerConfig ConfigFile
	PrivateIP    string
	ClientIP     string
}

func main() {

	var x ConfigFile

	// Read config from yaml file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	check(err)

	err = yaml.Unmarshal(yamlFile, &x)
	check(err)

	config := &AppConfig{
		DockerConfig: x,
		PrivateIP:    os.Args[2],
		ClientIP:     os.Args[2],
	}

	dockerCfg := &api.SwarmConfig{
		Endpoint:     config.DockerConfig.Endpoint,
		Nodetype:     config.DockerConfig.Nodetype,
		SwarmPort:    config.DockerConfig.SwarmPort,
		Managertoken: config.DockerConfig.ManagerToken,
		Workertoken:  config.DockerConfig.WorkerToken,
		PrivateIP:    config.PrivateIP,
		ClientIP:     config.ClientIP,
	}
	switch os.Args[1] {
	case "init":
		err = api.SwarmInit(*dockerCfg)
		check(err)

		token, err := api.SwarmTokens(*dockerCfg)
		check(err)

		z := &ConfigFile{
			Endpoint:     config.DockerConfig.Endpoint,
			Nodetype:     config.DockerConfig.Nodetype,
			SwarmPort:    config.DockerConfig.SwarmPort,
			ManagerToken: token["Manager"],
			WorkerToken:  token["Worker"],
		}

		yaml, err := yaml.Marshal(*z)
		check(err)

		err = ioutil.WriteFile("config.yaml", yaml, 0644)
		check(err)
	case "manager":
		resp, err := api.JoinSwarm(*dockerCfg)
		check(err)

		fmt.Println(resp)
	case "worker":
		fmt.Println("Worker")
	case "heal":
		fmt.Println("Healing")
	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
