package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

var (
	rc = `
_l() {
  local finalizer loon ret tmp
  loon="%s"

  tmp="$(mktemp -u)"
  exec 9>"${tmp}"
  exec 8<"${tmp}"
  rm ${tmp}

  "${loon}" "$@"
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

alias loon="_l"`

	rcCmd = &cli.Command{
		Name:   "rc",
		Hidden: true,
		Action: func(c *cli.Context) error {
			abs, _ := filepath.Abs(os.Args[0])
			rc := fmt.Sprintf(rc, abs)
			os.Stdout.Write([]byte(rc))
			return nil
		},
	}
)

func init() {
	appendCommand(rcCmd)
}
