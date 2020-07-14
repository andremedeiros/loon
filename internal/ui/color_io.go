package ui

import (
	"fmt"
	"io"
	"os"
)

func (c color) Write(p []byte) (n int, err error) {
	return c.Fprintf(os.Stdout, string(p))
}

func (color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fprintf(w, format, a...)
}

func fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	str := MustParse(fmt.Sprintf(format, a...), colorInstructionCodes)
	return fmt.Fprint(w, str)
}
