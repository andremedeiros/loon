package color

import (
	"fmt"
	"io"

	"github.com/andremedeiros/loon/internal/ui/parser"
)

func (color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	return fprintf(w, format, a...)
}

func fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	str := parser.MustParse(fmt.Sprintf(format, a...), InstructionCodes)
	return fmt.Fprint(w, str)
}
