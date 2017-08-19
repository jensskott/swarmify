package compute

import "github.com/rackspace/gophercloud/openstack/identity/v3/tokens"

func Connect(config tokens.AuthOptions) {
	scope := &tokens.Scope{ProjectName: "tmp_project"}

	token, err := tokens.Create(client, config, scope).Extract()
}
