package ui

import (
	"fmt"
	"io"
)

type SpinnerGroup interface {
	fmt.Stringer

	NewSpinner(string, ...interface{}) Spinner
	Update()
	Finish()
}

type Spinner interface {
	fmt.Stringer

	Do(func() error) error
	Update()
	Fail()
	Succeed()
}

type UI interface {
	NewSpinnerGroup() SpinnerGroup

	Info(string)
	Error(error)

	Fprintf(io.Writer, string, ...interface{}) (int, error)
}
