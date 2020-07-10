package cmd

import (
	"context"
	"flag"
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"
	"github.com/peterbourgon/usage"
)

var runEnv = func(ctx context.Context, ui ui.UI, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("env", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon env")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	return task.Run(ctx, ui, proj, "command:up", func(environ []string) error {
		sort.Strings(environ)
		w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
		for _, e := range environ {
			if strings.HasPrefix(e, "PATH=") {
				continue
			}
			es := strings.SplitN(e, "=", 2)
			ui.Fprintf(w, "{bold:%s}=%s\n", es[0], es[1])
		}
		w.Flush()
		return nil
	})
}
