package cmd

import (
	"context"
	"flag"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"
)

var runDown = func(ctx context.Context, ui ui.UI, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("down", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon down")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	return task.Run(ctx, ui, proj, "command:down", func(environ []string) error {
		ui.Info("All done!")
		return nil
	})
}
