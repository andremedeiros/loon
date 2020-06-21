package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	"github.com/andremedeiros/loon/internal/config"
)

func init() {
	rootCmd.AddCommand(shellRCCommand)
}

var shellRCCommand = &cobra.Command{
	Hidden: true,
	Use:    "shellrc",
	Short:  "Returns the shell initialization to use loon",
	Long:   `Returns the shell initialization to use loon`,
	Args:   cobra.ExactArgs(0),
	RunE: makeRunE(func(ctx context.Context, cfg *config.Config, cmd *cobra.Command, args []string) error {
		absPath, _ := filepath.Abs(os.Args[0])
		shell := strings.TrimSpace(fmt.Sprintf(`
__loon_path="%s"

_l() {
  local tmp ret finalizer

  tmp="$(mktemp -u)"
  exec 9>"${tmp}"
  exec 8<"${tmp}"
  rm ${tmp}

  "${__loon_path}" "$@"
  ret=$?

  while read -r finalizer; do
    case "${finalizer}" in
      chdir:*) cd "${finalizer//chdir:/}" ;;
      *) ;;
    esac
  done <&8

  exec 8<&-
  exec 9<&-

  return ${ret}
}
		`, absPath))

		cmd.OutOrStdout().Write([]byte(shell))
		return nil
	}),
}
