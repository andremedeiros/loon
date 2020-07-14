package cli

import (
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"

	"github.com/urfave/cli/v2"
)

var downCmd = &cli.Command{
	Name:    "down",
	Aliases: []string{"d", "land"},
	Usage:   "Shuts down project services",
	Action: func(c *cli.Context) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}
		cliui := ui.Instance()
		return task.Run(c.Context, cliui, proj, "command:down", func(environ []string) error {
			cliui.Info("All done!")
			return nil
		})
	},
}

func init() {
	appendCommand(downCmd)
}
