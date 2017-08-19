package api

// RunSwarmConfig run the swarm function
func RunSwarmConfig(config SwarmConfig) (string, error) {
	if config.Nodetype == "manager" {
		node, err := SearchSwarmCluster(config)
		if err != nil {
			return "", err
		}

		if node == nil {
			err = SwarmInit(config)
			if err != nil {
				return "", err
			}
			return "Docker swarm cluster initiated", nil
		}

		tkn, err := SwarmManagerToken(config)
		if err != nil {
			return "", err
		}

		return tkn, nil
	}

	return "", nil
}
