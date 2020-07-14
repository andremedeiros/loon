package cli

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/urfave/cli/v2"

	"github.com/andremedeiros/loon/internal/nix"
)

var versionsCmd = &cli.Command{
	Name:  "versions",
	Usage: "Lists versions of supported dependencies",
	Action: func(c *cli.Context) error {
		w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
		pkgs := nix.Packages()
		fmt.Fprintln(w, "Software\tVersion")
		for _, pkg := range pkgs {
			fmt.Fprintf(w, "%s\t%s\n", pkg.Name, pkg.Version)
		}
		w.Flush()
		return nil
	},
}

func init() {
	appendCommand(versionsCmd)
}
