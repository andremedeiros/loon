package cli

import (
	"errors"

	"github.com/andremedeiros/loon/internal/nix"
	"github.com/urfave/cli/v2"
)

var buildCmd = &cli.Command{
	Name:  "build",
	Usage: "Builds a dependency",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "all", Usage: "Builds all versions of a dependency"},
	},
	Action: func(c *cli.Context) error {
		all := c.Bool("all")
		if c.NArg() == 0 {
			return errors.New("specify a dependency name")
		}
		dep := c.Args().Get(0)

		if all {
			for _, pkg := range nix.Packages() {
				if pkg.Name != dep {
					continue
				}

				if err := pkg.Build(); err != nil {
					return err
				}
			}
		} else {
			ver := c.Args().Get(1)
			if ver == "" {
				ver = "default"
			}
			pkg, err := nix.PackageFor(dep, ver)
			if err != nil {
				return err
			}
			return pkg.Build()
		}
		return nil
	},
}

var nixCmd = &cli.Command{
	Name:        "nix",
	Usage:       "Nix toolbox",
	Subcommands: []*cli.Command{buildCmd},
}

func init() {
	appendCommand(nixCmd)
}
