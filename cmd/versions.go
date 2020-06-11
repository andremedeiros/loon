package cmd

import (
	"context"
	"fmt"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"github.com/andremedeiros/loon/catalog"
	"github.com/andremedeiros/loon/internal/config"
)

func init() {
	rootCmd.AddCommand(versionsCommand)
}

var versionsCommand = &cobra.Command{
	Use:   "versions",
	Short: "Prints all the available versions for the software that Loon handles",
	Long:  `Prints all the available versions for the software that Loon handles`,
	Args:  cobra.ExactArgs(0),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		w := tabwriter.NewWriter(cmd.OutOrStdout(), 0, 8, 4, '\t', 0)
		fmt.Fprintln(w, "Provider\tSoftware\tVersion")
		for _, sv := range catalog.List() {
			fmt.Fprintf(w, "%s\t%s\t%s\n", sv.Provider, sv.Software, sv.Version)
		}
		w.Flush()
		return nil
	}),
}
