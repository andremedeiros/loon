package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Info(msg string) {
	color.New(color.FgWhite, color.Bold).Fprint(os.Stdout, "Info: ")
	fmt.Fprintf(os.Stdout, "%s\n", msg)
}
