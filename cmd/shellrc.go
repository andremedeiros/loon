package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/peterbourgon/usage"

	"github.com/andremedeiros/loon/internal/config"
	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

var runShellRC = func(ctx context.Context, ui ui.UI, cfg *config.Config, _ *project.Project, args []string) error {
	flagset := flag.NewFlagSet("shellrc", flag.ExitOnError)
	flagset.Usage = usage.For(flagset, "loon shellrc")
	if err := flagset.Parse(args); err != nil {
		return err
	}

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

alias loon="_l"
	`, absPath))

	os.Stdout.Write([]byte(shell))
	return nil
}
