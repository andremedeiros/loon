package ui

import (
	"fmt"
	"io"
)

func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	str := MustParse(fmt.Sprintf(format, a...))
	return fmt.Fprint(w, str)
}
