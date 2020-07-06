package executor

import (
	"bufio"
	"bytes"
	"os/exec"
)

type Executor interface {
	Execute([]string, ...Option) error
}

func Execute(args []string, opts ...Option) error {
	stdout := bytes.Buffer{}
	stderr := bytes.Buffer{}
	{
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdout = bufio.NewWriter(&stdout)
		cmd.Stderr = bufio.NewWriter(&stderr)

		for _, opt := range opts {
			opt(cmd)
		}

		err := cmd.Run()
		code := cmd.ProcessState.ExitCode()
		if err != nil {
			err = NewExecutionError(err, args, code, stdout, stderr)
		}
		return err
	}
}
