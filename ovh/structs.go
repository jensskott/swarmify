package ovh

// Config for ovh
type Config struct {
	IdentityEndpoint string
	Username         string
	Password         string
	TenantID         string
	DomainName       string
	Region           string
	ImageID          string
	FlavorName       string
	Count            int
	Networks         []string
}

// PrivateIPAddress data
type PrivateIPAddress struct {
	OSEXTIPSMACMacAddr string `json:"OS-EXT-IPS-MAC:mac_addr"`
	OSEXTIPSType       string `json:"OS-EXT-IPS:type"`
	Addr               string `json:"addr"`
	Version            int    `json:"version"`
}

// IPAddress data
type IPAddress struct {
	Addr string `json:"addr"`
}
