package ovh

import (
	"github.com/rackspace/gophercloud/openstack/compute/v2/images"
)

// SearchImage by id
func SearchImage(config Config) (images.GetResult, error) {
	client, err := Connect(config)
	if err != nil {
		return images.GetResult{}, err
	}

	// opts := images.ListOpts{ChangesSince: "2014-01-01T01:02:03Z", Name: "Ubuntu 12.04"}
	return images.Get(client, "6c3a0a48-981a-4ccc-8d5e-e00c4dc3c4aa"), nil

}
