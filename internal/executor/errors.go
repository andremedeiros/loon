package executor

import (
	"bufio"
	"bytes"
	"strings"
)

type ExecutionError struct {
	err    error
	cmd    []string
	code   int
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewExecutionError(err error, cmd []string, code int, stdout bytes.Buffer, stderr bytes.Buffer) error {
	return ExecutionError{err, cmd, code, stdout, stderr}
}

func (e ExecutionError) Error() string {
	return e.err.Error()
}

func (e ExecutionError) Stdout() *bufio.Reader {
	return bufio.NewReader(&e.stdout)
}

func (e ExecutionError) Stderr() *bufio.Reader {
	return bufio.NewReader(&e.stderr)
}

func (e ExecutionError) Cmd() string {
	return strings.Join(e.cmd, " ")
}

func (e ExecutionError) Code() int {
	return e.code
}
