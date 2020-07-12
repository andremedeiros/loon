package executor

import (
	"bufio"
	"bytes"
	"os"
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
			if err := opt(cmd); err != nil {
				return err
			}
		}

		err := cmd.Run()
		code := cmd.ProcessState.ExitCode()
		if err != nil {
			err = NewExecutionError(err, args, code, stdout, stderr)
		}
		return err
	}
}

func RequestSudo(prompt string) error {
	cmd := exec.Command("sudo", "-p", prompt, "true")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
