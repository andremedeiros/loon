package cli

import (
	"os"
	"os/exec"

	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"

	"github.com/urfave/cli/v2"
)

var shellCmd = &cli.Command{
	Name:    "shell",
	Aliases: []string{"sh"},
	Usage:   "Starts a subshell with the project's environment",
	Action: func(c *cli.Context) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}
		cliui := ui.Instance()
		return task.Run(c.Context, cliui, proj, "command:up", func(environ []string) error {
			shell := os.Getenv("SHELL")
			cmd := exec.Command(shell)
			cmd.Env = append(os.Environ(), environ...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			code := cmd.ProcessState.ExitCode()
			os.Exit(code)
			return err
		})
	},
}

func init() {
	appendCommand(shellCmd)
}
