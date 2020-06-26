package executor

import (
	"io"
	"os/exec"
)

type Option func(*exec.Cmd)

func WithEnviron(environ []string) Option {
	return func(c *exec.Cmd) {
		c.Env = environ
	}
}

func WithStdout(w io.Writer) Option {
	return func(c *exec.Cmd) {
		c.Stdout = w
	}
}

func WithStderr(w io.Writer) Option {
	return func(c *exec.Cmd) {
		c.Stderr = w
	}
}

func WithStdin(r io.Reader) Option {
	return func(c *exec.Cmd) {
		c.Stdin = r
	}
}
