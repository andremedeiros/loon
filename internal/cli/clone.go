package cli

import (
	"errors"
	"os"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/finalizer"
	"github.com/andremedeiros/loon/internal/git"

	"github.com/urfave/cli/v2"
)

var cloneCmd = &cli.Command{
	Name:  "clone",
	Usage: "Clones a Git repository",
	Action: func(c *cli.Context) error {
		cfg, err := config.Read()
		if err != nil {
			return err
		}
		if c.NArg() == 0 {
			return errors.New("specify a partial project name")
		}
		partial := c.Args().Get(0)
		repo := git.NewRepository(partial)
		path := cfg.SourceTree.Resolve(repo.Host(), repo.Owner(), repo.Name())
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := repo.Clone(path)
			if err != nil {
				return err
			}
		}
		finalizer.Write("chdir", path)
		return nil
	},
}

func init() {
	appendCommand(cloneCmd)
}
