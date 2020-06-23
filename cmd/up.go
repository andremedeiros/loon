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
	fmt.Println("Ensuring software is installed...")
	if err = proj.EnsureDependencies(); err != nil {
		return err
	}
	fmt.Println("Setting up networking...")
	if err = proj.EnsureNetworking(); err != nil {
		return err
	}
	g, ctx := errgroup.WithContext(ctx)
	for _, srv := range proj.Services {
		srv := srv // otherwise it goes out of scope
		g.Go(func() error {
			if srv.IsHealthy(proj.IPAddr(), proj.VDPath()) {
				return nil
			}
			fmt.Printf("Starting %s...\n", srv.String())
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
					fmt.Printf("Something went wrong while initializing %s. Details:\n", srv.String())
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
					fmt.Printf("Something went wrong while starting %s. Details:\n", srv.String())
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
					return err
				}
			}
			return nil
		})
	}
	for _, lang := range proj.Languages {
		lang := lang // otherwise it goes out of scope
		fmt.Printf("Initializing %s...\n", lang.String())
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
				fmt.Printf("Something went wrong while initializing %s. Details:\n", lang.String())
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
