package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"text/tabwriter"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

type runHandler func(context.Context, *config.Config, *project.Project, []string) error

var version = "dev"

func rootUsage(p *project.Project) {
	cmd := filepath.Base(os.Args[0])
	ui.Fprintf(os.Stderr, "{bold:USAGE}\n")
	ui.Fprintf(os.Stderr, "  %s {cyan:<command>}\n", cmd)
	ui.Fprintf(os.Stderr, "\n")
	ui.Fprintf(os.Stderr, "{bold:COMMANDS}\n")
	ui.Fprintf(os.Stderr, "  {cyan:cd}        {yellow:Switches directories to project root}\n")
	ui.Fprintf(os.Stderr, "  {cyan:clone}     {yellow:Clones a Git repository into the working directory}\n")
	ui.Fprintf(os.Stderr, "  {cyan:doctor}    {yellow:Checks your system for potential problems}\n")
	ui.Fprintf(os.Stderr, "  {cyan:down}      {yellow:Stops the current project's infrastructure}\n")
	ui.Fprintf(os.Stderr, "  {cyan:exec}      {yellow:Executes command in project shell}\n")
	ui.Fprintf(os.Stderr, "  {cyan:shell}     {yellow:Starts a shell inheriting the current project's environment}\n")
	ui.Fprintf(os.Stderr, "  {cyan:up}        {yellow:Starts the current project's infrastructure}\n")
	ui.Fprintf(os.Stderr, "  {cyan:versions}  {yellow:Prints the versions of supported services and languages}\n")
	ui.Fprintf(os.Stderr, "\n")
	if p != nil && len(p.Tasks) > 0 {
		ui.Fprintf(os.Stderr, "{bold:PROJECT COMMANDS}\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		for _, t := range p.Tasks {
			ui.Fprintf(w, "\t{cyan:%s}\t{yellow:%s}\n", t.Name, t.Description)
		}
		w.Flush()
		ui.Fprintf(os.Stderr, "\n")
	}
	ui.Fprintf(os.Stderr, "{bold:VERSION}\n")
	ui.Fprintf(os.Stderr, "  %s (%s)\n", version, runtime.Version())
	ui.Fprintf(os.Stderr, "\n")
}

func Execute() error {
	proj, _ := project.FindInTree()
	if len(os.Args) < 2 {
		rootUsage(proj)
		os.Exit(1)
	}

	var run runHandler
	switch strings.ToLower(os.Args[1]) {
	case "cd":
		run = runCd
	case "clone":
		run = runClone
	case "doc", "doctor":
		run = runDoctor
	case "exec":
		run = runExec
	case "down", "land":
		run = runDown
	case "sh", "shell":
		run = runShell
	case "shellrc":
		run = runShellRC
	case "up", "fly":
		run = runUp
	case "versions":
		run = runVersions
	default:
		run = runTask(os.Args[1])
	}
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			s := <-sig
			fmt.Fprintf(os.Stderr, "\nExiting because of %q signal.\n", s)

			exitCode := 1
			if sysSig, ok := s.(syscall.Signal); ok {
				exitCode = int(sysSig)
			}

			cancel()
			os.Exit(exitCode + 128)
		}
	}()

	return run(ctx, cfg, proj, os.Args[2:])
}
