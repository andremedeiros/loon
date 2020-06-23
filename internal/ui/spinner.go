package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func Spinner(title string) (func(), func()) {
	s := spinner.New(
		spinner.CharSets[11],
		100*time.Millisecond,
		spinner.WithWriter(os.Stdout),
		spinner.WithColor("cyan"),
		spinner.WithSuffix(fmt.Sprintf(" %s", title)),
	)
	s.Start()

	failed := false

	success := func() {
		if failed {
			return
		}
		s.FinalMSG = fmt.Sprintf("%s %s\n", color.New(color.FgGreen).Sprintf("\u2713"), title)
		s.Stop()
	}
	failure := func() {
		failed = true
		s.FinalMSG = fmt.Sprintf("%s %s\n", color.New(color.FgRed).Sprintf("\u2717"), title)
		s.Stop()
	}
	return success, failure
}
