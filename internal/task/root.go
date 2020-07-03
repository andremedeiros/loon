package task

import (
	"fmt"
	"sync"

	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

type Root struct {
	Name string
}

func RunRoot(name string, p *project.Project) error {
	wg := &sync.WaitGroup{}
	wgs := make(map[string]*sync.WaitGroup)
	errs := make(chan error)
	ts, _ := graph.TopSort(name)
	// Do loop twice to ensure we're setup
	for _, t := range ts {
		wgs[t] = &sync.WaitGroup{}
		wgs[t].Add(1)
		wg.Add(1)
	}
	sg := ui.NewSpinnerGroup()
	for _, t := range ts {
		go func(t string) {
			defer func() {
				wgs[t].Done()
				wg.Done()
			}()
			te, ok := tasks[t]
			if !ok {
				return
			}
			// This waits for all the dependencies to be done.
			// It will last the longest amount of time a dependency
			// took to finish.
			for _, dep := range deps[t] {
				if dwg, ok := wgs[dep]; ok {
					dwg.Wait()
				}
			}
			done, err := te.Check(p)
			if err != nil {
				errs <- err
				return
			}
			if !done {
				s := sg.NewSpinner(te.Header())
				if err := te.Resolve(p); err != nil {
					s.Fail()
					errs <- err
				} else {
					done, err := te.Check(p)
					if err != nil {
						s.Fail()
						errs <- err
					} else if !done {
						s.Fail()
						errs <- fmt.Errorf("could not satisfy condition for %s", t)
					}
				}
				s.Succeed()
			}
		}(t)
	}
	go func() {
		wg.Wait()
		close(errs)
	}()
	for {
		select {
		case err := <-errs:
			return err
		}
	}
}
