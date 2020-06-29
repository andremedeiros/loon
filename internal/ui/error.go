package ui

import (
	"bytes"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
)

func Error(err error) {
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
		Fprintf(os.Stderr, "{red:{bold:Error:} Command invocation failed:}\n")
		Fprintf(os.Stderr, "%s\n", buf.String())
	default:
		Fprintf(os.Stderr, "{red:{bold:Error:} %s}\n", err)
	}
}
