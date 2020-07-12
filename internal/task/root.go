package task

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/andremedeiros/loon/internal/project"
	"github.com/andremedeiros/loon/internal/ui"
)

type Root struct {
	Name string
}

func Run(ctx context.Context, ui ui.UI, p *project.Project, name string, fun func([]string) error) error {
	bins := []string{}
	envs := []string{}
	wg := &sync.WaitGroup{}
	wgs := make(map[string]*sync.WaitGroup)
	errs := make(chan error)
	sg := ui.NewSpinnerGroup()
	bts, _ := graph.TopSort(name)
	ets, _ := graph.TopSort(name + ":done")
	// Do loop twice to ensure we're setup
	for _, ts := range [][]string{bts, ets} {
		for _, t := range ts {
			wgs[t] = &sync.WaitGroup{}
			wgs[t].Add(1)
			wg.Add(1)
		}
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
				done, err := te.Check(ctx, p)
				if err != nil {
					errs <- err
					return
				}
				e, b := te.Environ(ctx, p)
				envs = append(envs, e...)
				bins = append(bins, b...)
				if !done {
					s := sg.NewSpinner(te.Header())
					if err := te.Resolve(ctx, p); err != nil {
						s.Fail()
						errs <- err
						return
					} else {
						done, err := te.Check(ctx, p)
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
		// Wait for this particular group to be done. We want to wait for the
		// "before done" things to be finished before doing the "done" things
		// so that components like networking stop at the right time.
		wg.Wait()
	}
	// Set up environment
	sg.Finish()
	path := os.Getenv("PATH")
	envs = append(envs, fmt.Sprintf("PATH=%s:%s", strings.Join(bins, ":"), path))
	if err := fun(envs); err != nil {
		errs <- err
	}
	close(errs)
	for {
		select {
		case err := <-errs:
			sg.Finish()
			return err
		}
	}
}
