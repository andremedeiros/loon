package ui

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type SpinnerGroup struct {
	Spinners []*Spinner

	ticker    *time.Ticker
	refresh   bool
	lastShown int
	sync.Mutex
}

func NewSpinnerGroup() *SpinnerGroup {
	sg := &SpinnerGroup{ticker: time.NewTicker(100 * time.Millisecond)}
	go func() {
		for {
			select {
			case <-sg.ticker.C:
				sg.Update()
			}
		}
	}()
	return sg
}

func (sg *SpinnerGroup) NewSpinner(f string, a ...interface{}) *Spinner {
	label := fmt.Sprintf(f, a...)
	sg.Lock()
	s := NewSpinnerWithGroup(label, sg)
	sg.Spinners = append(sg.Spinners, s)
	sg.Unlock()
	sg.Update()
	return s
}

func (sg *SpinnerGroup) Update() {
	sg.Lock()
	defer sg.Unlock()
	Fprintf(os.Stdout, sg.String())
}

func (sg *SpinnerGroup) String() string {
	b := strings.Builder{}
	if sg.lastShown > 0 {
		Fprintf(os.Stdout, "\x1b[%dF", sg.lastShown)
	}
	sg.lastShown = 0
	for _, s := range sg.Spinners {
		Fprintf(&b, s.String())
		Fprintf(&b, "\n")
		sg.lastShown++
	}
	return b.String()
}
