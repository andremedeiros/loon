package cmd

import (
	"context"
	"flag"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
)

var runDown = func(ctx context.Context, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("down", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon down")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	return task.RunRoot("command:down", proj)
	/*
		sg := ui.NewSpinnerGroup()
		g, ctx := errgroup.WithContext(ctx)
		for _, srv := range proj.Services {
			srv := srv // otherwise it goes out of scope
			g.Go(func() error {
				s := sg.NewSpinner("Shutting down {cyan:%s}...", srv.String())
				return s.Do(func() error {
					if !srv.IsHealthy(proj.IP, proj.VDPath()) {
						return nil
					}
					if err := srv.Stop(proj, proj.IP, proj.VDPath()); err != nil {
						return err
					}
					return nil
				})
			})
		}
		return g.Wait()
	*/
}
