package ui

import (
	"fmt"
	"io"
	"os"
)

type simple struct{}

func NewSimple() simple {
	return simple{}
}

var simpleInstructionCodes = map[string]string{
	"reset":          "",
	"bold":           "",
	"faint":          "",
	"italic":         "",
	"underline":      "",
	"black":          "",
	"red":            "",
	"green":          "",
	"yellow":         "",
	"blue":           "",
	"magenta":        "",
	"cyan":           "",
	"white":          "",
	"bright_black":   "",
	"bright_red":     "",
	"bright_green":   "",
	"bright_yellow":  "",
	"bright_blue":    "",
	"bright_magenta": "",
	"bright_cyan":    "",
	"bright_white":   "",
}

func (s simple) Write(p []byte) (n int, err error) {
	return s.Fprintf(os.Stdout, string(p))
}

func (simple) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	str := MustParse(fmt.Sprintf(format, a...), simpleInstructionCodes)
	return fmt.Fprintf(w, str)
}

func (simple) Error(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
}

func (simple) Info(msg string) {
	fmt.Fprintf(os.Stdout, "%s\n", msg)
}

type spinnergroup struct{}

func (simple) NewSpinnerGroup() SpinnerGroup {
	return &spinnergroup{}
}

func (sg *spinnergroup) NewSpinner(f string, a ...interface{}) Spinner {
	label := fmt.Sprintf(f, a...)
	fmt.Fprintf(os.Stdout, "%s\n", label)
	return &spinner{label}
}

func (sg *spinnergroup) String() string { return "" }
func (sg *spinnergroup) Update()        {}
func (sg *spinnergroup) Pause()         {}
func (sg *spinnergroup) Resume()        {}
func (sg *spinnergroup) Finish()        {}
func (sg *spinnergroup) ResetDisplay()  {}

type spinner struct {
	label string
}

func (s *spinner) String() string            { return s.label }
func (s *spinner) Do(fun func() error) error { return fun() }
func (s *spinner) Update()                   {}
func (s *spinner) Fail()                     {}
func (s *spinner) Wait()                     {}
func (s *spinner) Work()                     {}
func (s *spinner) Succeed()                  {}
