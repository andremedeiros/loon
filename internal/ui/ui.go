package ui

import (
	"fmt"
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

type SpinnerGroup interface {
	fmt.Stringer

	NewSpinner(string, ...interface{}) Spinner
	Update()
	Pause()
	Resume()
	Finish()
	ResetDisplay()
}

type Spinner interface {
	fmt.Stringer

	Do(func() error) error
	Update()
	Fail()
	Wait()
	Work()
	Succeed()
}

type UI interface {
	NewSpinnerGroup() SpinnerGroup

	Info(string)
	Error(error)

	Fprintf(io.Writer, string, ...interface{}) (int, error)
}

var _ UI = NewColor()
var _ UI = NewSimple()

func Instance() UI {
	if IsTTY() {
		return NewColor()
	} else {
		return NewSimple()
	}
}

func IsTTY() bool {
	return isatty.IsTerminal(os.Stdout.Fd())
}
