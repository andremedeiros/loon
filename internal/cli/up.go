package cli

import (
	"github.com/urfave/cli/v2"

	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"
)

var upCmd = &cli.Command{
	Name:    "up",
	Aliases: []string{"u", "fly"},
	Usage:   "Starts up project services",
	Action: func(c *cli.Context) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}
		cliui := ui.Instance()
		return task.Run(c.Context, cliui, proj, "command:up", func(environ []string) error {
			cliui.Info("Ready to go!")
			return nil
		})
	},
}

func init() {
	appendCommand(upCmd)
}
