package process

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func FromPidFile(path string) (*os.Process, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pid, _ := strconv.Atoi(strings.TrimSpace(string(b)))
	return os.FindProcess(pid)
}

func InterruptFromPidFile(path string) error {
	proc, err := FromPidFile(path)
	if err != nil {
		return err
	}
	return proc.Signal(os.Interrupt)
}
