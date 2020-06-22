package cmd

import (
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/usage"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
)

var runDown = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("down", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon down")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	proj, err := project.FindInTree()
	if err != nil {
		return err
	}

	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range proj.Services {
		fmt.Printf("stopping %s...\n", srv.String())
		srv := srv // otherwise it goes out of scope
		g.Go(func() error {
			return srv.Stop(proj, proj.IPAddr(), proj.VDPath())
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
