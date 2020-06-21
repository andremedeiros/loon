package cmd

import (
	"context"
	"flag"
	"os"
	"os/exec"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/peterbourgon/usage"
)

var runShell = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("shell", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon shell")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	proj, err := project.FindInTree()
	if err != nil {
		return err
	}

	shell := os.Getenv("SHELL")
	ex := exec.Command(shell)
	ex.Env = proj.Environ()
	ex.Stdout = os.Stdout
	ex.Stdin = os.Stdin
	ex.Stderr = os.Stderr
	return ex.Run()
}
