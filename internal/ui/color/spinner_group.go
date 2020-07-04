package color

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/andremedeiros/loon/internal/ui"
)

type SpinnerGroup struct {
	Spinners []*Spinner

	ticker    *time.Ticker
	refresh   bool
	lastShown int
	c         color
	sync.Mutex
}

func (c color) NewSpinnerGroup() ui.SpinnerGroup {
	sg := &SpinnerGroup{ticker: time.NewTicker(100 * time.Millisecond), c: c}
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

func (sg *SpinnerGroup) NewSpinner(f string, a ...interface{}) ui.Spinner {
	label := fmt.Sprintf(f, a...)
	sg.Lock()
	s := sg.c.NewSpinnerWithGroup(label, sg)
	sg.Spinners = append(sg.Spinners, s)
	sg.Unlock()
	sg.Update()
	return s
}

func (sg *SpinnerGroup) Update() {
	sg.Lock()
	defer sg.Unlock()
	sg.c.Fprintf(os.Stdout, sg.String())
}

func (sg *SpinnerGroup) Finish() {
	sg.ticker.Stop()
}

func (sg *SpinnerGroup) String() string {
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
