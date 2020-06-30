package cmd

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

var runExec = func(ctx context.Context, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("exec", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon exec <cmd> [arg1] [arg2]")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if flagset.NArg() < 1 {
		return errors.New("specify a command")
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	if proj.NeedsUpdate() {
		return errors.New("project needs update, run `loon up`")
	}
	if err := proj.Execute(
		flagset.Args(),
		executor.WithStdin(os.Stdin),
		executor.WithStdout(os.Stdout),
		executor.WithStderr(os.Stderr),
	); err != nil {
		err := err.(executor.ExecutionError)
		os.Exit(err.Code())
		return err
	}
	return nil
}
