package main

import (
	"fmt"
	"log"

	api "github.com/jensskott/swarmify/api"
)

func main() {

	config := &api.SwarmConfig{
		Endpoint: "unix:///var/run/docker.sock",
		Nodetype: "manager",
	}

	resp, err := api.RunSwarmConfig(*config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
