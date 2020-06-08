package cmd

import (
	"context"
	"fmt"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/git"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pathCommand)
}

var pathCommand = &cobra.Command{
	Use:   "path [repository name with owner]",
	Short: "Prints the project path for a specific repository",
	Long: `Prints the project path for a specific repository

Some accepted values are:

  $ loon path andremedeiros/loon
	$ loon path andre.cool`,
	Args: cobra.ExactArgs(1),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		repo := git.NewRepository(args[0])
		path := cfg.SourceTree.Resolve(repo.Host(), repo.Owner(), repo.Name())
		fmt.Println(path)
		return nil
	}),
}
