# My project's README

    /*
	ovhCfg := &ovh.Config{
		IdentityEndpoint: config.OvhConfig.IdentityEndpoint,
		Username:         config.OvhConfig.Username,
		Password:         config.OvhConfig.Password,
		TenantID:         config.OvhConfig.TenantID,
		DomainName:       config.OvhConfig.DomainName,
		Region:           config.OvhConfig.Region,
		ImageID:          config.OvhConfig.ImageID,
		FlavorName:       config.OvhConfig.FlavorName,
		Count:            os.Args[2],
	}

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
    */
    
    
    	client, err := Connect(config)
	if err != nil {
		return nil, err
	}

	network1 := servers.Network{UUID: config.Networks[0]}
	network2 := servers.Network{UUID: config.Networks[1]}

	networks := []servers.Network{network1, network2}

	name := ("swarm" + nodetype)

	serverOpts := &servers.CreateOpts{
		Name:       name,
		FlavorName: config.FlavorName,
		ImageRef:   config.ImageID,
		Networks:   networks,
	}

	server, err := servers.Create(client, *serverOpts).Extract()
	if err != nil {
		return nil, err
	}

	ips, err := waitForIP(client, server.ID)
	if err != nil {
		return nil, err
	}

	// Just when testing
	servers.Delete(client, server.ID)

	return ips, nil


func waitForIP(client *gophercloud.ServiceClient, id string) ([]string, error) {

	var ips []string
	nets := []string{"VLAN-Static", "Ext-Net"}

	for _, i := range nets {
		var ipData []IPAddress

		for {
			srv, _ := servers.Get(client, id).Extract()
			if srv.Addresses[i] == nil {
				time.Sleep(time.Second * 20)
				continue
			}

			jsonByte, err := json.Marshal(srv.Addresses[i])
			if err != nil {
				return nil, err
			}

			err = json.Unmarshal(jsonByte, &ipData)
			if err != nil {
				return nil, err
			}

			fmt.Println(srv.Addresses[i])

			ips = append(ips, ipData[0].Addr)

			if len(ips) > 2 {
				break
			}
		}
		return removeDuplicates(ips), nil
	}

	return nil, nil
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}

	for v := range elements {
		if encountered[elements[v]] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[elements[v]] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}



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