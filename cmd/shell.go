package cmd

import (
	"context"
	"errors"
	"flag"
	"os"
	"os/exec"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
)

var runShell = func(ctx context.Context, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("shell", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon shell")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	if proj.NeedsUpdate() {
		return errors.New("project needs update, run `loon up`")
	}
	shell := os.Getenv("SHELL")
	cmd := exec.Command(shell)
	cmd.Env = proj.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	code := cmd.ProcessState.ExitCode()
	os.Exit(code)
	return err
}
