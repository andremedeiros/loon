package check

import (
	"errors"
	"os/exec"
	"time"
)

var ErrSudoNotEnabled = errors.New("sudo not enabled for this user")

func Sudo() error {
	cmd := exec.Command("sudo", "true")
	errs := make(chan error)
	go func() {
		err := cmd.Run()
		code := cmd.ProcessState.ExitCode()
		if err != nil || code != 0 {
			errs <- ErrSudoNotEnabled
		} else {
			errs <- nil
		}
	}()

	timeout := time.After(2 * time.Second)
	select {
	case <-timeout:
		cmd.Process.Kill()
		return ErrSudoNotEnabled
	case err := <-errs:
		return err
	}
}
