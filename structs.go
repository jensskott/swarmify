package main

// DockerConfigFile for the run
type DockerConfigFile struct {
	Endpoint     string `yaml:"endpoint"`
	Nodetype     string `yaml:"nodetype"`
	SwarmPort    string `yaml:"swarmport"`
	WorkerToken  string `yaml:"workertoken"`
	ManagerToken string `yaml:"managertoken"`
}

// AppConfig for the whole app
type AppConfig struct {
	DockerConfig DockerConfigFile
	OvhConfig    OvhConfigFile
	PrivateIP    string
	ClientIP     string
}

// OvhConfigFile for the run
type OvhConfigFile struct {
	Identityendpoint string `json:"identityendpoint"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Tenantid         string `json:"tenantid"`
	Tenantname       string `json:"tenantname"`
	Domainname       string `json:"domainname"`
	Region           string `json:"region"`
	Imageid          string `json:"imageid"`
	Flavorname       string `json:"flavorname"`
	Rules            []struct {
		Name     string `json:"name"`
		Protocol string `json:"protocol"`
		Fromport int    `json:"fromport"`
		Toport   int    `json:"toport"`
		Cidr     string `json:"cidr"`
	} `json:"rules"`
}
