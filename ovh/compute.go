package ovh

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// CreateCompute resource for ovh
func CreateCompute(dir string) (map[string]string, error) {

	init := exec.Command("terraform", "init")
	init.Dir = dir
	init.Run()

	cmd := exec.Command("terraform", "apply")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(fmt.Sprint(err) + ": " + string(output))
	}

	ips, err := getIps(dir)
	if err != nil {
		return nil, err
	}

	return ips, nil
}

func getIps(dir string) (map[string]string, error) {
	var lines []string
	m := make(map[string]string)

	cmd := exec.Command("terraform", "output")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errors.New(fmt.Sprint(err) + ": " + string(output))
	}

	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, i := range lines {
		x := strings.Split(i, "=")
		y := strings.Split(x[1], ":")
		m[y[0]] = y[1]
	}

	return m, nil
}
