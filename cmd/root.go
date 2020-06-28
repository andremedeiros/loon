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
)

type runHandler func(context.Context, *config.Config, []string) error

var version = "dev"

func rootUsage(p *project.Project) {
	cmd := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "USAGE\n")
	fmt.Fprintf(os.Stderr, "  %s <command>\n", cmd)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "COMMANDS\n")
	fmt.Fprintf(os.Stderr, "  cd        Switches directories to project root\n")
	fmt.Fprintf(os.Stderr, "  clone     Clones a Git repository into the working directory\n")
	fmt.Fprintf(os.Stderr, "  doctor    Checks your system for potential problems\n")
	fmt.Fprintf(os.Stderr, "  down      Stops the current project's infrastructure\n")
	fmt.Fprintf(os.Stderr, "  exec      Executes command in project shell\n")
	fmt.Fprintf(os.Stderr, "  shell     Starts a shell inheriting the current project's environment\n")
	fmt.Fprintf(os.Stderr, "  up        Starts the current project's infrastructure\n")
	fmt.Fprintf(os.Stderr, "  versions  Prints the versions of supported services and languages\n")
	fmt.Fprintf(os.Stderr, "\n")
	if p != nil && len(p.Tasks) > 0 {
		fmt.Fprintf(os.Stderr, "PROJECT COMMANDS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		for _, t := range p.Tasks {
			fmt.Fprintf(w, "\t%s\t%s\n", t.Name, t.Description)
		}
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
	fmt.Fprintf(os.Stderr, "VERSION\n")
	fmt.Fprintf(os.Stderr, "  %s (%s)\n", version, runtime.Version())
	fmt.Fprintf(os.Stderr, "\n")
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

	return run(ctx, cfg, os.Args[2:])
}
