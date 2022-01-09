package cli

import "os/exec"

func RunCommand(context string) error {
	return exec.Command("kubectl", "config", "use-context", context).Run()
}
