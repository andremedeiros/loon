package check

import (
	"errors"
	"os/exec"
	"time"
)

func Sudo() error {
	cmd := exec.Command("sudo", "true")
	cmd.Start()

	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	timeout := time.After(2 * time.Second)
	select {
	case <-timeout:
		cmd.Process.Kill()
		return errors.New("sudo not enabled in this system")
	case <-done:
		return nil
	}
}
