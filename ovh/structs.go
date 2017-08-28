package ovh

// Config for ovh
type Config struct {
	IdentityEndpoint string
	Username         string
	Password         string
	TenantID         string
	TenantName       string
	DomainName       string
	Region           string
	ImageID          string
	FlavorName       string
	Count            string
}

// PrivateIPAddress data
type PrivateIPAddress struct {
	OSEXTIPSMACMacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	OSEXTIPSType       string `json:"OS-EXT-IPS:type"`
	Addr               string `json:"addr"`
	Version            int    `json:"version"`
}
