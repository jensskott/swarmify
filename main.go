package main

import (
	"fmt"
	"log"

	api "github.com/jensskott/swarmify/api"
)

func main() {

	nodeType := "manager"
	endPoint := "unix:///var/run/docker.sock"

	resp, err := api.RunSwarmConfig(nodeType, endPoint)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}
