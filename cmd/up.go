package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/provider"
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

		g, ctx := errgroup.WithContext(ctx)
		for _, h := range provider.Handlers {
			fmt.Printf("installing with %s...\n", h.String())
			h := h // otherwise it goes out of scope
			g.Go(h.Install)
		}
		if err := g.Wait(); err != nil {
			return err
		}

		g, ctx = errgroup.WithContext(ctx)
		for _, srv := range proj.Services {
			fmt.Printf("starting %s...\n", srv.String())
			srv := srv // otherwise it goes out of scope
			g.Go(srv.Start)
		}
		if err := g.Wait(); err != nil {
			return err
		}

		return nil
	}),
}
