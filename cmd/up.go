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

var runUp = func(ctx context.Context, ui ui.UI, cfg *config.Config, proj *project.Project, args []string) error {
	flagset := flag.NewFlagSet("up", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon up")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	if proj == nil {
		return project.ErrProjectPayloadNotFound
	}
	return task.Run(ctx, ui, proj, "command:up", func(environ []string) error {
		ui.Info("Ready to go!")
		return nil
	})
}
