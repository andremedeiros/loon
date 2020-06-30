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

var runUp = func(ctx context.Context, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("up", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon up")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	sg := ui.NewSpinnerGroup()
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		if !proj.NeedsUpdate() {
			return nil
		}
		s := sg.NewSpinner("Installing dependencies...")
		return s.Do(proj.EnsureDependencies)
	})

	g.Go(func() error {
		if !proj.NeedsNetworking() {
			return nil
		}
		s := sg.NewSpinner("Setting up networking...")
		return s.Do(proj.EnsureNetworking)
	})

	if err := g.Wait(); err != nil {
		return err
	}

	for _, srv := range proj.Services {
		srv := srv // otherwise it goes out of scope
		g.Go(func() error {
			s := sg.NewSpinner("Starting {cyan:%s}...", srv.String())
			return s.Do(func() error {
				if srv.IsHealthy(proj.IP, proj.VDPath()) {
					return nil
				}
				if err := srv.Initialize(proj, proj.IP, proj.VDPath()); err != nil {
					return err
				}
				if err := srv.Start(proj, proj.IP, proj.VDPath()); err != nil {
					return err
				}
				return nil
			})
		})
	}

	for _, lang := range proj.Languages {
		lang := lang // otherwise it goes out of scope
		g.Go(func() error {
			s := sg.NewSpinner("Setting up {cyan:%s}...", lang.String())
			return s.Do(func() error {
				return lang.Initialize(proj, proj.VDPath())
			})
		})
	}
	return g.Wait()
}
