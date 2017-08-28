package ovh

import (
	"bufio"
	"os/exec"
	"strings"
)

// CreateCompute resource for ovh
func CreateCompute(nodetype string) (map[string]string, error) {

	cmd := exec.Command("terraform", "apply")
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return getIps(), nil
}

func getIps() map[string]string {
	var lines []string
	m := make(map[string]string)

	cmd, _ := exec.Command("terraform", "output").Output()

	scanner := bufio.NewScanner(strings.NewReader(string(cmd)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, i := range lines {
		x := strings.Split(i, "=")
		y := strings.Split(x[1], ":")
		m[y[0]] = y[1]
	}

	return m
}
