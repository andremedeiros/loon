package cli

import (
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"
)

func taskCommand(p *project.Project, t *project.Task) func(*cli.Context) error {
	return func(c *cli.Context) error {
		cliui := ui.Instance()
		return task.Run(c.Context, cliui, p, "command:up", func(environ []string) error {
			tmp, err := ioutil.TempFile("", "loon.sh")
			if err != nil {
				return err
			}
			os.Chmod(tmp.Name(), 0770)
			defer os.Remove(tmp.Name())
			tmp.Write([]byte(t.Command))
			tmp.Close()
			if err := p.Execute(
				append([]string{tmp.Name()}, c.Args().Slice()...),
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

func init() {
	p, err := project.FindInTree()
	if err != nil {
		return
	}
	for _, t := range p.Tasks {
		t := t
		tc := &cli.Command{
			Name:     t.Name,
			Usage:    t.Description,
			Category: "Project Tasks",
			Action:   taskCommand(p, &t),
		}
		appendCommand(tc)
	}
}
