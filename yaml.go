package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

func readYaml(dockerYaml, ovhYaml string) (DockerConfigFile, OvhConfigFile) {
	var x DockerConfigFile
	var y OvhConfigFile

	dockerYamlFile, err := ioutil.ReadFile(dockerYaml)
	check(err)

	ovhYamlFile, err := ioutil.ReadFile(ovhYaml)
	check(err)

	err = yaml.Unmarshal(dockerYamlFile, &x)
	check(err)

	err = yaml.Unmarshal(ovhYamlFile, &y)

	return x, y
}

func writeYaml(dockerStruct *DockerConfigFile, dockerYaml string) {
	yaml, err := yaml.Marshal(*dockerStruct)
	check(err)

	err = ioutil.WriteFile(dockerYaml, yaml, 0644)
	check(err)
}
