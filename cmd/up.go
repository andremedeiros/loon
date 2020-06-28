package cmd

import (
	"context"
	"flag"

	"github.com/peterbourgon/usage"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

var runUp = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("up", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon up")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	proj, err := project.FindInTree()
	if err != nil {
		return err
	}

	if proj.NeedsUpdate() {
		success, failure := ui.Spinner("Ensuring software is installed...")
		if err = proj.EnsureDependencies(); err != nil {
			failure()
			return err
		}
		success()
	}

	if proj.NeedsNetworking() {
		success, failure := ui.Spinner("Setting up networking...")
		if err := proj.EnsureNetworking(); err != nil {
			failure()
			return err
		}
		success()
	}

	success, failure := ui.Spinner("Starting...")
	defer success()
	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range proj.Services {
		srv := srv // otherwise it goes out of scope
		g.Go(func() error {
			if srv.IsHealthy(proj.IP, proj.VDPath()) {
				return nil
			}
			if err := srv.Initialize(proj, proj.IP, proj.VDPath()); err != nil {
				failure()
				return err
			}
			if err := srv.Start(proj, proj.IP, proj.VDPath()); err != nil {
				failure()
				return err
			}
			return nil
		})
	}

	for _, lang := range proj.Languages {
		lang := lang // otherwise it goes out of scope
		g.Go(func() error {
			err := lang.Initialize(proj, proj.VDPath())
			if err != nil {
				failure()
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	success()
	return nil
}
