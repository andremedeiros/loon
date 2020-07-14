package cli

import (
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/task"
	"github.com/andremedeiros/loon/internal/ui"

	"github.com/urfave/cli/v2"
)

var envCmd = &cli.Command{
	Name:  "env",
	Usage: "Prints the project's environment",
	Action: func(c *cli.Context) error {
		proj, err := project.FindInTree()
		if err != nil {
			return err
		}
		cliui := ui.Instance()
		return task.Run(c.Context, cliui, proj, "command:up", func(environ []string) error {
			sort.Strings(environ)
			w := tabwriter.NewWriter(os.Stdout, 0, 2, 2, ' ', 0)
			for _, e := range environ {
				switch {
				case strings.HasPrefix(e, "LOON_NEW_ENVS="):
				case strings.HasPrefix(e, "LOON_OLD_ENV_"):
				case strings.HasPrefix(e, "LOON_PROJECT_ROOT="):
				case strings.HasPrefix(e, "PATH="):
				default:
					es := strings.SplitN(e, "=", 2)
					cliui.Fprintf(w, "{bold:%s}=%s\n", es[0], es[1])
				}
			}
			w.Flush()
			return nil
		})
	},
}

func init() {
	appendCommand(envCmd)
}
