package cmd

import (
	"context"
	"flag"
	"io/ioutil"
	"os"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"
)

var runTask = func(taskName string) runHandler {
	return func(ctx context.Context, ui ui.UI, cfg *config.Config, proj *project.Project, args []string) error {
		flagset := flag.NewFlagSet(taskName, flag.ContinueOnError)
		flagset.Usage = usage.For(flagset, "loon <task>")
		flagset.Parse(args)
		if proj == nil {
			return project.ErrProjectPayloadNotFound
		}
		t, err := proj.Task(taskName)
		if err != nil {
			return err
		}
		return task.Run(ctx, ui, proj, "command:up", func(environ []string) error {
			tmp, err := ioutil.TempFile("", "loon.sh")
			if err != nil {
				return err
			}
			os.Chmod(tmp.Name(), 0770)
			defer os.Remove(tmp.Name())
			tmp.Write([]byte(t.Command))
			if err := proj.Execute(
				append([]string{tmp.Name()}, flagset.Args()...),
				executor.WithStdin(os.Stdin),
				executor.WithStdout(os.Stdout),
				executor.WithStderr(os.Stderr),
				executor.WithEnviron(environ),
			); err != nil {
				err := err.(executor.ExecutionError)
				os.Exit(err.Code())
				return err
			}
			return nil
		})
	}
}
