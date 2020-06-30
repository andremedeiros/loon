package cmd

import (
	"context"
	"errors"
	"flag"
	"os"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/finalizer"
	"github.com/andremedeiros/loon/internal/git"
	"github.com/andremedeiros/loon/internal/project"
)

var runClone = func(ctx context.Context, cfg *config.Config, _ *project.Project, args []string) error {
	flagset := flag.NewFlagSet("clone", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon clone <owner and name>")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	args = flagset.Args()
	if flagset.NArg() < 1 {
		return errors.New("specify an owner and name")
	}
	repo := git.NewRepository(flagset.Arg(0))
	path := cfg.SourceTree.Resolve(repo.Host(), repo.Owner(), repo.Name())
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := repo.Clone(path)
		if err != nil {
			return err
		}
	}
	finalizer.Write("chdir", path)
	return nil
}
