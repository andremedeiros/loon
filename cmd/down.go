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
		srv := srv // otherwise it goes out of scope
		g.Go(func() error {
			if !srv.IsHealthy(proj.IPAddr(), proj.VDPath()) {
				return nil
			}
			fmt.Printf("Stopping %s...\n", srv.String())
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
				fmt.Printf("Something went wrong while stopping %s. Details:\n", srv.String())
				if stdout.Len() > 0 {
					fmt.Println("-------------------- 8< stdout 8< --------------------")
					fmt.Println(stdout.String())
					fmt.Println("------------------------------------------------------")
				}
				if stderr.Len() > 0 {
					fmt.Println("-------------------- 8< stderr 8< --------------------")
					fmt.Println(stderr.String())
					fmt.Println("------------------------------------------------------")
				}
			}
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
