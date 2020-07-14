package cli

import (
	"context"

	"github.com/urfave/cli/v2"
)

var version = "dev"

var app = &cli.App{
	Name:  "Loon",
	Usage: "Your development swiss knife",
	Authors: []*cli.Author{
		{Name: "André Medeiros", Email: "hello@andre.cool"},
	},
	Version:              version,
	EnableBashCompletion: true,
}

func RunContext(ctx context.Context, args []string) error {
	return app.RunContext(ctx, args)
}

func appendCommand(c *cli.Command) {
	app.Commands = append(app.Commands, c)
}
