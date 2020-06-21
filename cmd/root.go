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

	"github.com/andremedeiros/loon/internal/config"
)

type runHandler func(context.Context, *config.Config, []string) error

var version = "dev"

func rootUsage() {
	cmd := filepath.Base(os.Args[0])
	fmt.Fprintf(os.Stderr, "USAGE\n")
	fmt.Fprintf(os.Stderr, "  %s <command>\n", cmd)
	fmt.Fprintf(os.Stderr, "\n")
	fmt.Fprintf(os.Stderr, "COMMANDS\n")
	fmt.Fprintf(os.Stderr, "  cd        Switches directories to project root\n")
	fmt.Fprintf(os.Stderr, "  clone     Clones a Git repository into the working directory\n")
	fmt.Fprintf(os.Stderr, "  down      Stops the current project's infrastructure\n")
	fmt.Fprintf(os.Stderr, "  shell     Starts a shell inheriting the current project's environment\n")
	fmt.Fprintf(os.Stderr, "  up        Starts the current project's infrastructure\n")
	fmt.Fprintf(os.Stderr, "  versions  Prints the versions of supported services and languages\n")
	fmt.Fprintf(os.Stderr, "VERSION\n")
	fmt.Fprintf(os.Stderr, "  %s (%s)\n", version, runtime.Version())
	fmt.Fprintf(os.Stderr, "\n")
}

func Execute() error {
	if len(os.Args) < 2 {
		rootUsage()
		os.Exit(1)
	}

	var run runHandler
	switch strings.ToLower(os.Args[1]) {
	case "cd":
		run = runCd
	case "clone":
		run = runClone
	case "down":
	case "land":
		run = runDown
	case "shell":
		run = runShell
	case "shellrc":
		run = runShellRC
	case "up":
	case "fly":
		run = runUp
	case "versions":
		run = runVersions
	default:
		rootUsage()
		os.Exit(1)
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

	return run(ctx, cfg, os.Args[1:])
}
