package color

import (
	"fmt"
	"os"
)

var spinnerTheme = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

type SpinnerState int

const (
	Working SpinnerState = iota
	Waiting
	Succeeded
	Failed
)

func (ss SpinnerState) String(s *Spinner) string {
	switch ss {
	case Succeeded:
		return "{green:\u2713}"
	case Failed:
		return "{red:\u2717}"
	case Waiting:
		return "{yellow:\u203C}"
	default:
		s.curIdx = (s.curIdx + 1) % len(spinnerTheme)
		return fmt.Sprintf("{cyan:%s}", spinnerTheme[s.curIdx])
	}
}

type Spinner struct {
	Label string
	Theme []string
	State SpinnerState

	sg     *SpinnerGroup
	curIdx int
	c      color
}

func (c color) NewSpinner(label string) *Spinner {
	return c.NewSpinnerWithGroup(label, nil)
}

func (c color) NewSpinnerWithGroup(label string, sg *SpinnerGroup) *Spinner {
	return &Spinner{label, spinnerTheme, Working, sg, 0, c}
}

func (s *Spinner) Do(fun func() error) error {
	if err := fun(); err != nil {
		s.Fail()
		return err
	}
	s.Succeed()
	return nil
}

func (s *Spinner) Update() {
	if s.sg != nil {
		s.sg.Update()
		return
	}
	s.c.Fprintf(os.Stdout, s.String())
}

func (s *Spinner) Fail() {
	s.State = Failed
	s.Update()
}

func (s *Spinner) Succeed() {
	s.State = Succeeded
	s.Update()
}

func (s *Spinner) Wait() {
	if s.State == Working {
		s.State = Waiting
	}
}

func (s *Spinner) Work() {
	if s.State == Waiting {
		s.State = Working
	}
}

func (s *Spinner) String() string {
	// Figure out how many cycles
	return fmt.Sprintf("\r\x1b[2K%s %s", s.State.String(s), s.Label)
}
