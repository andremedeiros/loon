package color

import (
	"os"
)

func (c color) Info(msg string) {
	c.Fprintf(os.Stdout, "{blue:%s}\n", msg)
}
