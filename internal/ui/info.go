package ui

import (
	"os"
)

func Info(msg string) {
	Fprintf(os.Stdout, "{blue:{bold,underline:Info:} %s}\n", msg)
}
