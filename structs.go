package main

// DockerConfigFile for the run
type DockerConfigFile struct {
	Endpoint     string `yaml:"endpoint"`
	Nodetype     string `yaml:"nodetype"`
	SwarmPort    string `yaml:"swarmport"`
	WorkerToken  string `yaml:"workertoken"`
	ManagerToken string `yaml:"managertoken"`
}

// OvhConfigFile for the run
type OvhConfigFile struct {
	IdentityEndpoint string `yaml:"identityendpoint"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	TenantID         string `yaml:"tenantid"`
	TenantName       string `yaml:"tenantname"`
	DomainName       string `yaml:"domainname"`
	Region           string `yaml:"region"`
	ImageID          string `yaml:"imageid"`
	FlavorName       string `yaml:"flavorname"`
}

// AppConfig for the whole app
type AppConfig struct {
	DockerConfig DockerConfigFile
	OvhConfig    OvhConfigFile
	PrivateIP    string
	ClientIP     string
}
