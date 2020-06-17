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
	rootCmd.AddCommand(downCommand)
}

var downCommand = &cobra.Command{
	Use:   "down",
	Short: "STOPS ALL THE THINGS",
	Long:  `Stops all the things`,
	Args:  cobra.ExactArgs(0),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
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
	}),
}
