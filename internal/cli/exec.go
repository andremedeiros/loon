package cli

import (
	"errors"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"

	"github.com/urfave/cli/v2"
)

var execCmd = &cli.Command{
	Name:    "exec",
	Aliases: []string{"e"},
	Usage:   "Executes a command in the project's environment",
	Action: func(c *cli.Context) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}
		if c.NArg() == 0 {
			return errors.New("specify a command")
		}
		cliui := ui.Instance()
		return task.Run(c.Context, cliui, proj, "command:up", func(environ []string) error {
			if err := proj.Execute(
				c.Args().Slice(),
				executor.WithStdin(os.Stdin),
				executor.WithStdout(os.Stdout),
				executor.WithStderr(os.Stderr),
				executor.WithEnviron(environ),
			); err != nil {
				err := err.(executor.ExecutionError)
				os.Exit(err.Code())
				return err
			}
			return nil
		})
	},
}

func init() {
	appendCommand(execCmd)
}
