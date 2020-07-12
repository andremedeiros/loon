package executor

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

type Option func(*exec.Cmd) error

func WithEnv(env map[string]string) Option {
	environ := []string{}
	for k, v := range env {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}
	return WithEnviron(environ)
}

func WithEnviron(environ []string) Option {
	return func(c *exec.Cmd) error {
		c.Env = append(os.Environ(), environ...)
		return nil
	}
}

func WithStdout(w io.Writer) Option {
	return func(c *exec.Cmd) error {
		c.Stdout = w
		return nil
	}
}

func WithStderr(w io.Writer) Option {
	return func(c *exec.Cmd) error {
		c.Stderr = w
		return nil
	}
}

func WithStdin(r io.Reader) Option {
	return func(c *exec.Cmd) error {
		c.Stdin = r
		return nil
	}
}

func WithSudo(fun func() error) Option {
	return func(c *exec.Cmd) error {
		cmd := exec.Command("sudo", "-n", "true")
		c.Args = append([]string{c.Path}, c.Args...)
		c.Path = cmd.Path
		// Check whether we can sudo without password.
		if cmd.Run() == nil {
			return nil
		}
		return fun()
	}
}
