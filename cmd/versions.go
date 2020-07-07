package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/nix"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

var runVersions = func(ctx context.Context, ui ui.UI, cfg *config.Config, _ *project.Project, args []string) error {
	flagset := flag.NewFlagSet("versions", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon versions")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	fmt.Fprintln(w, "Software\tVersion")
	for _, pkg := range nix.Packages() {
		fmt.Fprintf(w, "%s\t%s\n", pkg.Name, pkg.Version)
	}
	w.Flush()
	return nil
}
