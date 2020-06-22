package cmd

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/executer"
	"github.com/andremedeiros/loon/internal/project"
)

var runTask = func(ctx context.Context, cfg *config.Config, args []string) error {
	taskName := args[0]
	flagset := flag.NewFlagSet(taskName, flag.ExitOnError)
	flagset.Usage = usage.For(flagset, fmt.Sprintf("loon %s", taskName))
	if err := flagset.Parse(args); err != nil {
		return err
	}
	proj, err := project.FindInTree()
	if err != nil {
		return err
	}
	for _, task := range proj.Tasks {
		if task.Name != taskName {
			continue
		}
		tmp, err := ioutil.TempFile("", "loon.sh")
		if err != nil {
			return err
		}
		os.Chmod(tmp.Name(), 0770)
		defer os.Remove(tmp.Name())
		tmp.Write([]byte(task.Command))
		code, err := proj.Execute(
			[]string{tmp.Name()},
			executer.WithStdin(os.Stdin),
			executer.WithStdout(os.Stdout),
			executer.WithStderr(os.Stderr),
		)
		os.Exit(code)
		return err
	}
	return fmt.Errorf("task not found: %s", taskName)
}
