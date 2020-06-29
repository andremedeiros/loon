package ui

import (
	"os"
)

func Info(msg string) {
	Fprintf(os.Stdout, "{blue:{bold:Info:} %s}\n", msg)
}
