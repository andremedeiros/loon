package ui

import (
	"fmt"
	"os"
)

var spinnerTheme = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

type spinnerState int

const (
	Working spinnerState = iota
	Waiting
	Succeeded
	Failed
)

func (ss spinnerState) String(s *colorSpinner) string {
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

type colorSpinner struct {
	Label string
	Theme []string
	State spinnerState

	sg     *colorSpinnerGroup
	curIdx int
	c      color
}

func (c color) NewSpinner(label string) *colorSpinner {
	return c.NewSpinnerWithGroup(label, nil)
}

func (c color) NewSpinnerWithGroup(label string, sg *colorSpinnerGroup) *colorSpinner {
	return &colorSpinner{label, spinnerTheme, Working, sg, 0, c}
}

func (s *colorSpinner) Do(fun func() error) error {
	if err := fun(); err != nil {
		s.Fail()
		return err
	}
	s.Succeed()
	return nil
}

func (s *colorSpinner) Update() {
	if s.sg != nil {
		s.sg.Update()
		return
	}
	s.c.Fprintf(os.Stdout, s.String())
}

func (s *colorSpinner) Fail() {
	s.State = Failed
	s.Update()
}

func (s *colorSpinner) Succeed() {
	s.State = Succeeded
	s.Update()
}

func (s *colorSpinner) Wait() {
	if s.State == Working {
		s.State = Waiting
	}
}

func (s *colorSpinner) Work() {
	if s.State == Waiting {
		s.State = Working
	}
}

func (s *colorSpinner) String() string {
	return fmt.Sprintf("\r\x1b[2K%s %s", s.State.String(s), s.Label)
}
