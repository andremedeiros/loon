package check

import (
	"errors"
	"os/exec"
)

var ErrSudoNotEnabled = errors.New("sudo not enabled for this user")

func Sudo() error {
	cmd := exec.Command("sudo", "-n", "true")
	err := cmd.Run()
	code := cmd.ProcessState.ExitCode()
	if err != nil || code != 0 {
		return ErrSudoNotEnabled
	}
	return nil
}
