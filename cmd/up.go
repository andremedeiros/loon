package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
)

func init() {
	rootCmd.AddCommand(upCommand)
}

var upCommand = &cobra.Command{
	Use:   "up",
	Short: "INSTALLS ALL THE THINGS",
	Long:  `Installs all the things`,
	Args:  cobra.ExactArgs(0),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}

		fmt.Println("ensuring software installed...")
		if err = proj.EnsureDependencies(); err != nil {
			return err
		}

		fmt.Println("setting up networking...")
		if err = proj.EnsureNetworking(); err != nil {
			return err
		}

		g, ctx := errgroup.WithContext(ctx)
		for _, srv := range proj.Services {
			srv := srv // otherwise it goes out of scope
			g.Go(func() error {
				fmt.Printf("starting %s...\n", srv.String())
				if err := srv.Initialize(proj, proj.IPAddr(), proj.VDPath()); err != nil {
					return err
				}
				return srv.Start(proj, proj.IPAddr(), proj.VDPath())
			})
		}
		for _, lang := range proj.Languages {
			lang := lang // otherwise it goes out of scope
			g.Go(func() error {
				fmt.Printf("setting up %s...\n", lang.String())
				return lang.Initialize(proj, proj.VDPath())
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}

		return nil
	}),
}
