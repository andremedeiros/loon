package cmd

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"fmt"

	"github.com/peterbourgon/usage"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/executer"
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

	{
		success, failure := ui.Spinner("Ensuring software is installed...")
		if err = proj.EnsureDependencies(); err != nil {
			failure()
			return err
		}
		success()
	}

	{
		success, failure := ui.Spinner("Setting up networking...")
		if err = proj.EnsureNetworking(); err != nil {
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
			if srv.IsHealthy(proj.IPAddr(), proj.VDPath()) {
				return nil
			}
			// Initialize
			{
				stdout := bytes.Buffer{}
				stderr := bytes.Buffer{}
				err := srv.Initialize(
					proj,
					proj.IPAddr(),
					proj.VDPath(),
					executer.WithStdout(bufio.NewWriter(&stdout)),
					executer.WithStderr(bufio.NewWriter(&stderr)),
				)
				if err != nil {
					failure()
					ui.ErrorWithOutput(
						fmt.Sprintf("Something went wrong while initializing %s", srv.String()),
						stdout,
						stderr,
					)
					return err
				}
			}

			// Start
			{
				stdout := bytes.Buffer{}
				stderr := bytes.Buffer{}
				err := srv.Start(
					proj,
					proj.IPAddr(),
					proj.VDPath(),
					executer.WithStdout(bufio.NewWriter(&stdout)),
					executer.WithStderr(bufio.NewWriter(&stderr)),
				)
				if err != nil {
					failure()
					ui.ErrorWithOutput(
						fmt.Sprintf("Something went wrong while starting %s", srv.String()),
						stdout,
						stderr,
					)
					return err
				}
			}
			return nil
		})
	}

	for _, lang := range proj.Languages {
		lang := lang // otherwise it goes out of scope
		g.Go(func() error {
			stdout := bytes.Buffer{}
			stderr := bytes.Buffer{}
			err := lang.Initialize(
				proj,
				proj.VDPath(),
				executer.WithStdout(bufio.NewWriter(&stdout)),
				executer.WithStderr(bufio.NewWriter(&stderr)),
			)
			if err != nil {
				failure()
				ui.ErrorWithOutput(
					fmt.Sprintf("Something went wrong while initializing %s", lang.String()),
					stdout,
					stderr,
				)
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
