package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/finalizer"
	"github.com/andremedeiros/loon/internal/git"
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

var cloneCmd = &cobra.Command{
	Use:   "clone [repository name with owner]",
	Short: "Clones a Git repository into the work directory",
	Long: `Clone will resolve a repository identifier and clone it to the disk.

Some accepted values are:

  $ loon clone andremedeiros/loon
	$ loon clone andre.cool
	$ loon clone git@github.com:andremedeiros/k6.git`,
	Args: cobra.ExactArgs(1),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		repo := git.NewRepository(args[0])
		path := cfg.SourceTree.Resolve(repo.Host(), repo.Owner(), repo.Name())
		err := repo.Clone(path)
		if err != nil {
			return err
		}
		finalizer.Write("chdir", path)
		return nil
	}),
}
