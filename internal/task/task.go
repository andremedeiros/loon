package task

import (
	"fmt"
	"net"
	"time"

	"github.com/andremedeiros/loon/internal/project"

	"github.com/stevenle/topsort"
)

var (
	graph                         = topsort.NewGraph()
	deps  map[string][]string     = make(map[string][]string)
	tasks map[string]TaskExecutor = make(map[string]TaskExecutor)
)

type TaskExecutor interface {
	Header() string
	Check(*project.Project) (bool, error)
	Resolve(*project.Project) error
}

func RegisterTask(name string, t TaskExecutor) {
	if _, ok := tasks[name]; ok {
		panic(name)
	}
	tasks[name] = t
}

func Depends(what string, on string) {
	if _, ok := deps[what]; !ok {
		deps[what] = []string{}
	}
	deps[what] = append(deps[what], on)
	graph.AddNode(what)
	graph.AddNode(on)
	graph.AddEdge(on, what)
}

func Task(name string) TaskExecutor {
	if task, ok := tasks[name]; ok {
		return task
	}
	panic(name)
}

func checkHealth(ip net.IP, port int) bool {
	done := make(chan bool)
	go func() {
		hp := fmt.Sprintf("%s:%d", ip, port)
		for {
			if _, err := net.Dial("tcp", hp); err == nil {
				done <- true
			}
		}
	}()
	select {
	case <-done:
		return true
	case <-time.After(200 * time.Millisecond):
		return false
	}
}

func checkProjectHasLanguage(p *project.Project, l string) bool {
	for _, lang := range p.Languages {
		if l == lang.String() {
			return true
		}
	}
	return false
}
func checkProjectHasService(p *project.Project, s string) bool {
	for _, srv := range p.Services {
		if s == srv.String() {
			return true
		}
	}
	return false
}
