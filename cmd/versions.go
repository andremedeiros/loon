package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/catalog"
	"github.com/andremedeiros/loon/internal/config"
)

var runVersions = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("versions", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon versions")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 8, 4, '\t', 0)
	fmt.Fprintln(w, "Software\tVersion")
	for _, sv := range catalog.List() {
		fmt.Fprintf(w, "%s\t%s\n", sv.Name, sv.Version)
	}
	w.Flush()
	return nil
}
