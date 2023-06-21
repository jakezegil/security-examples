package utilities

import (
	"fmt"
	"os/exec"
	"strings"
)

func RouterIP() (gateway string) {
	gateway, err := findDefaultGateway()

	if err == nil {
		fmt.Println("Default gateway:", gateway)
		return gateway
	}

	fmt.Println("Failed to find the gateway IP address.")
	return
}

func findDefaultGateway() (string, error) {
	cmd := exec.Command("netstat", "-nr")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	fmt.Println(string(output))

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		// in some cases, this may be 0.0.0.0 ?
		if len(fields) > 1 && fields[0] == "default" {
			return fields[1], nil
		}
	}

	return "", fmt.Errorf("gateway IP address not found")
}
