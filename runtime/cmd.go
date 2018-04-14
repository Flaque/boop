package runtime

import (
	"bytes"
	"os/exec"
)

func RunCmd(name string, args ...string) string {

	// Attempt to call if from the command line
	cmd := exec.Command(name, args...)

	//TODO Set stdin

	// Set stdout
	var out bytes.Buffer
	cmd.Stdout = &out

	// Launch command
	cmd.Run()

	return out.String()
}
