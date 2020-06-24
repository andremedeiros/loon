package cmd

import (
	"context"
	"errors"
	"flag"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/finalizer"
	"github.com/andremedeiros/loon/internal/git"
)

var runCd = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("cd", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon cd <project>")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	args = flagset.Args()
	if len(args) <= 0 {
		return errors.New("specify a partial project name")
	}
	repo := git.NewRepository(args[1])
	path := cfg.SourceTree.Resolve(repo.Host(), repo.Owner(), repo.Name())
	finalizer.Write("chdir", path)
	return nil
}
