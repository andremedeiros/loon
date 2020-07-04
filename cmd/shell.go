package cmd

import (
	"context"
	"flag"
	"os"
	"os/exec"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"
)

var runShell = func(ctx context.Context, ui ui.UI, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("shell", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon shell")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	return task.Run(ctx, ui, proj, "command:up", func(environ []string) error {
		shell := os.Getenv("SHELL")
		cmd := exec.Command(shell)
		cmd.Env = append(os.Environ(), environ...)
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		code := cmd.ProcessState.ExitCode()
		os.Exit(code)
		return err
	})
}
