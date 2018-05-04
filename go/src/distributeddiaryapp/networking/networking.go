package networking

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetOutboundIP Returns a machine's public (outbound) Azure IP address e.g. "270.0.21.1"
// We make the assumption that all production apps will be run on an Azure VM.
func GetOutboundIP() (ipString string, err error) {
	out, err := exec.Command("curl", "-s", "http://checkip.amazonaws.com").Output()

	if err != nil {
		fmt.Println(err)
	}

	result := fmt.Sprintf("%s", out)
	return strings.TrimSpace(result), nil
}
