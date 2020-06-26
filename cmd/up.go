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
		stdout := bytes.Buffer{}
		stderr := bytes.Buffer{}
		err = proj.EnsureNetworking(
			executer.WithStdout(bufio.NewWriter(&stdout)),
			executer.WithStderr(bufio.NewWriter(&stderr)),
		)
		if err != nil {
			failure()
			ui.ErrorWithOutput(
				"Something went wrong while setting up network",
				stdout,
				stderr,
			)
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
			// Initialize
			{
				stdout := bytes.Buffer{}
				stderr := bytes.Buffer{}
				err := srv.Initialize(
					proj,
					proj.IP,
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
					proj.IP,
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
