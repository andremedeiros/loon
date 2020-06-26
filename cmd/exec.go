package cmd

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/executer"
	"github.com/andremedeiros/loon/internal/project"
)

var runExec = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("exec", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon exec <cmd> [arg1] [arg2]")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if flagset.NArg() < 1 {
		return errors.New("specify a command")
	}
	proj, err := project.FindInTree()
	if err != nil {
		return err
	}
	if proj.NeedsUpdate() {
		return errors.New("project needs update, run `loon up`")
	}
	code, err := proj.Execute(
		flagset.Args(),
		executer.WithStdin(os.Stdin),
		executer.WithStdout(os.Stdout),
		executer.WithStderr(os.Stderr),
	)
	os.Exit(code)
	return err
}
