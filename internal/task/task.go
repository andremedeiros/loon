package task

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/andremedeiros/loon/internal/project"

	"github.com/stevenle/topsort"
)

var (
	graph = topsort.NewGraph()
	deps  = make(map[string][]string)
	tasks = make(map[string]Executor)
)

type Executor interface {
	Header() string
	Check(context.Context, *project.Project) (bool, error)
	Resolve(context.Context, *project.Project) error
	Environ(context.Context, *project.Project) (Environ, BinPaths)
}

type Environ []string
type BinPaths []string

func RegisterTask(name string, t Executor) {
	if _, ok := tasks[name]; ok {
		panic(name)
	}
	tasks[name] = t
}

func RunsAfter(on string, what string) {
	if _, ok := deps[what]; !ok {
		deps[what] = []string{}
	}
	deps[what] = append(deps[what], on)
	graph.AddNode(what)
	graph.AddNode(on)
	graph.AddEdge(on, what)
}

func checkHealth(ip net.IP, port int, waitUp bool) bool {
	done := make(chan bool)
	go func() {
		hp := fmt.Sprintf("%s:%d", ip, port)
		for {
			conn, err := net.Dial("tcp", hp)
			if conn != nil {
				conn.Close()
			}
			isUp := (err == nil)
			if !isUp && waitUp {
				continue
			}
			done <- isUp
		}
	}()
	select {
	case healthy := <-done:
		return healthy
	case <-time.After(5000 * time.Millisecond):
		return false
	}
}

func checkProjectHasDep(p *project.Project, d string) bool {
	for _, dep := range p.Dependencies {
		if d == dep.Name {
			return true
		}
	}
	return false
}
