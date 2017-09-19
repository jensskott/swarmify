package ovh

import (
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/secgroups"
)

func CreateInfra(config Config) (string, error) {
	client, err := Connect(config)
	if err != nil {
		return "", err
	}

	groupOpts := secgroups.CreateOpts{
		Name:        "docker_sg",
		Description: "for docker communication",
	}

	group, err := secgroups.Create(client, groupOpts).Extract()
	if err != nil {
		return "", err
	}

	for _, r := range config.Rules {
		ruleOpts := secgroups.CreateRuleOpts{
			ParentGroupID: group.ID,
			FromPort:      r.Fromport,
			ToPort:        r.Toport,
			IPProtocol:    r.Protocol,
			CIDR:          r.Cidr,
		}

		_, err := secgroups.CreateRule(client, ruleOpts).Extract()
		if err != nil {
			return "", err
		}

	}

	ruleOptsTcp := secgroups.CreateRuleOpts{
		ParentGroupID: group.ID,
		FromPort:      0,
		ToPort:        65535,
		IPProtocol:    "tcp",
		CIDR:          group.Name,
	}

	err = secgroups.CreateRule(client, ruleOptsTcp).Extract()
	if err != nil {
		return "", err
	}

	ruleOptsUdp := secgroups.CreateRuleOpts{
		ParentGroupID: group.ID,
		FromPort:      0,
		ToPort:        65535,
		IPProtocol:    "udp",
		CIDR:          group.Name,
	}

	err = secgroups.CreateRule(client, ruleOptsUdp).Extract()
	if err != nil {
		return "", err
	}

	return "Security groups created", nil
}
