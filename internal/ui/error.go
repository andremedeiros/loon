package ui

import (
	"bytes"
	"fmt"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/fatih/color"
)

func Error(err error) {
	msg := err.Error()
	switch err := err.(type) {
	case executor.ExecutionError:
		buf := bytes.Buffer{}
		if stdout := err.Stdout(); stdout.Size() > 0 {
			buf.WriteString("-------------------- 8< stdout 8< --------------------\n")
			stdout.WriteTo(&buf)
		}
		if stderr := err.Stderr(); stderr.Size() > 0 {
			buf.WriteString("-------------------- 8< stderr 8< --------------------\n")
			stderr.WriteTo(&buf)
		}
		msg = fmt.Sprintf("command invocation failed.\n%s", buf.String())
	}
	color.New(color.FgRed, color.Bold).Fprint(os.Stderr, "Error: ")
	fmt.Fprintf(os.Stderr, "%s\n", msg)
}
