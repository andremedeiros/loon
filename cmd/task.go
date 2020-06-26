package cmd

import (
	"context"
	"errors"
	"flag"
	"io/ioutil"
	"os"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/executer"
	"github.com/andremedeiros/loon/internal/project"
)

var runTask = func(taskName string) runHandler {
	return func(ctx context.Context, cfg *config.Config, args []string) error {
		flagset := flag.NewFlagSet(taskName, flag.ContinueOnError)
		flagset.Usage = usage.For(flagset, "loon <task>")
		flagset.Parse(args)
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}
		if proj.NeedsUpdate() {
			return errors.New("project needs update, run `loon up`")
		}
		task, err := proj.Task(taskName)
		if err != nil {
			return err
		}
		tmp, err := ioutil.TempFile("", "loon.sh")
		if err != nil {
			return err
		}
		os.Chmod(tmp.Name(), 0770)
		defer os.Remove(tmp.Name())
		tmp.Write([]byte(task.Command))
		code, err := proj.Execute(
			append([]string{tmp.Name()}, flagset.Args()...),
			executer.WithStdin(os.Stdin),
			executer.WithStdout(os.Stdout),
			executer.WithStderr(os.Stderr),
		)
		os.Exit(code)
		return err
	}
}
