package cmd

import (
	"context"
	"flag"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/peterbourgon/usage"
)

var runDoctor = func(ctx context.Context, cfg *config.Config, args []string) error {
	flagset := flag.NewFlagSet("doctor", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon doctor")
	if err := flagset.Parse(args); err != nil {
		return err
	}
	// TODO:
	// check nix is installed
	// check sudo is setup
	// check for available diskspace
	return nil
}
