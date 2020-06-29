package executor

import (
	"bufio"
	"bytes"
	"os/exec"
)

type Executor interface {
	Execute([]string, ...Option) (int, error)
}

func Execute(cmd []string, opts ...Option) (int, error) {
	name := cmd[0]
	args := cmd[1:]
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	{
		cmd := exec.Command(name, args...)
		cmd.Stdout = bufio.NewWriter(&stdout)
		cmd.Stderr = bufio.NewWriter(&stderr)

		for _, opt := range opts {
			opt(cmd)
		}

		err := cmd.Run()
		code := cmd.ProcessState.ExitCode()
		if err != nil {
			err = NewExecutionError(err, stdout, stderr)
		}
		return code, err
	}
}
