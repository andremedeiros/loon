package executor

import (
	"bufio"
	"bytes"
)

type ExecutionError struct {
	err    error
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewExecutionError(err error, stdout bytes.Buffer, stderr bytes.Buffer) error {
	return ExecutionError{err, stdout, stderr}
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
