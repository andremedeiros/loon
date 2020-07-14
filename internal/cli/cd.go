package cli

import (
	"errors"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/finalizer"
	"github.com/andremedeiros/loon/internal/git"

	"github.com/urfave/cli/v2"
)

var cdCmd = &cli.Command{
	Name:  "cd",
	Usage: "Switches directories inside the sources root",
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
		finalizer.Write("chdir", path)
		return nil
	},
}

func init() {
	appendCommand(cdCmd)
}
