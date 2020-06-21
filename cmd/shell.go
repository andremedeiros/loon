package cmd

import (
	"context"
	"os"
	"os/exec"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(shellCommand)
}

var shellCommand = &cobra.Command{
	Use:     "shell",
	Aliases: []string{"sh"},
	Short:   "Starts a shell with the project's environment setup",
	Long:    `Starts a shell with the project's environment setup`,
	Args:    cobra.ExactArgs(0),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}

		shell := os.Getenv("SHELL")
		ex := exec.Command(shell)
		ex.Env = proj.Environ()
		ex.Stdout = cmd.OutOrStdout()
		ex.Stdin = cmd.InOrStdin()
		ex.Stderr = cmd.ErrOrStderr()
		return ex.Run()
	}),
}
