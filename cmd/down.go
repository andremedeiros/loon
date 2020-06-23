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
	success, failure := ui.Spinner("Stopping...")
	defer success()
	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range proj.Services {
		srv := srv // otherwise it goes out of scope
		g.Go(func() error {
			if !srv.IsHealthy(proj.IPAddr(), proj.VDPath()) {
				return nil
			}
			stdout := bytes.Buffer{}
			stderr := bytes.Buffer{}
			err := srv.Stop(
				proj,
				proj.IPAddr(),
				proj.VDPath(),
				executer.WithStdout(bufio.NewWriter(&stdout)),
				executer.WithStderr(bufio.NewWriter(&stderr)),
			)
			if err != nil {
				failure()
				ui.ErrorWithOutput(
					fmt.Sprintf("Something went wrong while stopping %s", srv.String()),
					stdout,
					stderr,
				)
			}
			return err
		})
	}
	return g.Wait()
}
