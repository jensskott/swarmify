package api

// SwarmConfig for the api
type SwarmConfig struct {
	Endpoint     string
	Nodetype     string
	Managertoken string
	Workertoken  string
	SwarmMaster  string
	SwarmPort    string
	PrivateIP    string
	ClientIP     string
}
