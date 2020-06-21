package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/finalizer"
	"github.com/andremedeiros/loon/internal/git"
)

func init() {
	rootCmd.AddCommand(cdCommand)
}

var cdCommand = &cobra.Command{
	Use:   "cd [repository name with owner]",
	Short: "Switches directories to the project path for a specific repository",
	Long: `Switches directories to the project path for a specific repository

Some accepted values are:

  $ loon cd andremedeiros/loon
	$ loon cd andre.cool`,
	Args: cobra.ExactArgs(1),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		repo := git.NewRepository(args[0])
		path := cfg.SourceTree.Resolve(repo.Host(), repo.Owner(), repo.Name())
		finalizer.Write("chdir", path)
		return nil
	}),
}
