package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	bashAutocomplete = `
		_loon_autocomplete() {
			if [[ "${COMP_WORDS[0]}" != "source" ]]; then
				local cur opts base
				COMPREPLY=()
				cur="${COMP_WORDS[COMP_CWORD]}"
				if [[ "$cur" == "-"* ]]; then
					opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
				else
					opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
				fi
				COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
				return 0
			fi
		}

		complete -o bashdefault -o default -o nospace -F _loon_autocomplete _l
	`

	zshAutocomplete = `
		#compdef _l

		_loon_autocomplete() {
			local -a opts
			local cur
			cur=${words[-1]}
			if [[ "$cur" == "-"* ]]; then
				opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
			else
				opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
			fi

			if [[ "${opts[1]}" != "" ]]; then
				_describe 'values' opts
			else
				_files
			fi

			return
		}

		compdef _loon_autocomplete _l
	`

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

		alias loon="_l"
	`
)
var rcCmd = &cli.Command{
	Name:   "rc",
	Hidden: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "bash", Usage: "Generate .bashrc snippet"},
		&cli.BoolFlag{Name: "zsh", Usage: "Generate .zshrc snippet"},
	},
	Action: func(c *cli.Context) error {
		bash := c.Bool("bash")
		zsh := c.Bool("zsh")
		if bash == zsh {
			return errors.New("you have to pick either --bash or --zsh")
		}
		buf := strings.Builder{}
		abs, _ := filepath.Abs(os.Args[0])
		if bash {
			buf.WriteString(stripIndent(bashAutocomplete))
		}
		if zsh {
			buf.WriteString(stripIndent(zshAutocomplete))
		}
		rc := fmt.Sprintf(rc, abs)
		buf.WriteString(stripIndent(rc))
		fmt.Println(buf.String())
		return nil
	},
}

func stripIndent(s string) string {
	re := regexp.MustCompile(`(?m)^\t\t`)
	return string(re.ReplaceAll([]byte(s), []byte("")))
}

func init() {
	appendCommand(rcCmd)
}
