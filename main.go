package main

import (
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

func main() {

	var config ConfigFile

	// Read config from yaml file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	check(err)

	err = yaml.Unmarshal(yamlFile, &config)
	check(err)

	dockerCfg := &api.SwarmConfig{
		Endpoint:  config.Endpoint,
		Nodetype:  config.Nodetype,
		SwarmPort: config.SwarmPort,
		PrivateIP: "127.0.0.1",
		ClientIP:  "127.0.0.1",
	}

	if os.Args[1] == "init" {
		err = api.SwarmInit(*dockerCfg)
		check(err)

		token, err := api.SwarmTokens(*dockerCfg)
		check(err)

		cfg := &ConfigFile{
			Endpoint:     config.Endpoint,
			Nodetype:     config.Nodetype,
			SwarmPort:    config.SwarmPort,
			WorkerToken:  token["Worker"],
			ManagerToken: token["Manager"],
		}

		yaml, err := yaml.Marshal(*cfg)
		check(err)

		err = ioutil.WriteFile("config.yaml", yaml, 0644)
		check(err)

	}
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
