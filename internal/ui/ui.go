package ui

import (
	"fmt"
	"io"
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
