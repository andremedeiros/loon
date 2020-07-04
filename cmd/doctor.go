package cmd

import (
	"context"
	"flag"

	"github.com/peterbourgon/usage"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/check"
	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

var checks = []func() error{
	check.Sudo,
	check.Nix,
}

var runDoctor = func(ctx context.Context, ui ui.UI, cfg *config.Config, _ *project.Project, args []string) error {
	flagset := flag.NewFlagSet("doctor", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon doctor")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, check := range checks {
		check := check
		g.Go(func() error {
			if err := check(); err != nil {
				ui.Error(err)
				return err
			}
			return nil
		})
	}
	if g.Wait() == nil {
		ui.Info("You're all good!")
	}
	return nil
}
