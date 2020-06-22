package cmd

import (
	"context"
	"errors"
	"flag"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
)

var runExec = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("exec", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon exec <cmd> [arg1] [arg2]")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	args = flagset.Args()
	if len(args) < 2 {
		return errors.New("specify a command")
	}
	proj, err := project.FindInTree()
	if err != nil {
		return err
	}
	return proj.Execute(args[1:])
}
