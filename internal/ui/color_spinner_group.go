package ui

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

type colorSpinnerGroup struct {
	Spinners []*colorSpinner

	ticker    *time.Ticker
	refresh   bool
	paused    bool
	lastShown int
	c         color
	sync.Mutex
}

func (c color) NewSpinnerGroup() SpinnerGroup {
	sg := &colorSpinnerGroup{ticker: time.NewTicker(100 * time.Millisecond), c: c}
	go func() {
		for {
			select {
			case <-sg.ticker.C:
				if !sg.paused {
					sg.Update()
				}
			}
		}
	}()
	return sg
}

func (sg *colorSpinnerGroup) NewSpinner(f string, a ...interface{}) Spinner {
	label := fmt.Sprintf(f, a...)
	sg.Lock()
	s := sg.c.NewSpinnerWithGroup(label, sg)
	sg.Spinners = append(sg.Spinners, s)
	sg.Unlock()
	sg.Update()
	return s
}

func (sg *colorSpinnerGroup) Update() {
	sg.Lock()
	defer sg.Unlock()
	if sg.paused {
		return
	}
	sg.c.Fprintf(os.Stdout, sg.String())
}

func (sg *colorSpinnerGroup) Finish() {
	sg.ticker.Stop()
}

func (sg *colorSpinnerGroup) Pause() {
	sg.Update()
	sg.paused = true
}

func (sg *colorSpinnerGroup) Resume() {
	sg.paused = false
}

func (sg *colorSpinnerGroup) ResetDisplay() {
	sg.lastShown = 0
}

func (sg *colorSpinnerGroup) String() string {
	b := strings.Builder{}
	if sg.lastShown > 0 {
		sg.c.Fprintf(os.Stdout, "\x1b[%dF", sg.lastShown)
	}
	sg.lastShown = 0
	for _, s := range sg.Spinners {
		sg.c.Fprintf(&b, s.String())
		sg.c.Fprintf(&b, "\n")
		sg.lastShown++
	}
	return b.String()
}
