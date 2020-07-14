package cli

import (
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"

	"github.com/andremedeiros/loon/internal/check"
	"github.com/andremedeiros/loon/internal/ui"
)

var checks = []func() error{
	check.Nix,
}

var doctorCmd = &cli.Command{
	Name:    "doctor",
	Aliases: []string{"doc"},
	Usage:   "Checks the system for potential problems",
	Action: func(c *cli.Context) error {
		cliui := ui.Instance()
		g, _ := errgroup.WithContext(c.Context)
		for _, check := range checks {
			check := check
			g.Go(func() error {
				if err := check(); err != nil {
					cliui.Error(err)
					return err
				}
				return nil
			})
		}
		if g.Wait() == nil {
			cliui.Info("You're all good!")
		}
		return nil
	},
}

func init() {
	appendCommand(doctorCmd)
}
