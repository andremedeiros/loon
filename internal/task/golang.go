package task

import (
	"context"
	"fmt"

	"github.com/andremedeiros/loon/internal/project"
)

func goEnviron(p *project.Project) []string {
	golang := p.VariableDataPath("data", "go")
	return []string{
		fmt.Sprintf("GOPATH=%s", golang),
	}
}

type GoInitialize struct{}

func (*GoInitialize) Header() string {
	return "Setting up {blue:Golang}"
}

func (*GoInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	return true, nil
}

func (*GoInitialize) Resolve(_ context.Context, p *project.Project) error {
	return nil
}

func (*GoInitialize) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	if !checkProjectHasDep(p, "golang") {
		return nil, nil
	}
	bin := p.VariableDataPath("data", "go", "bin")
	return goEnviron(p), []string{bin}
}

func init() {
	RegisterTask("golang:initialize", &GoInitialize{})
	RunsAfter("derivation:current:up", "golang:initialize")
}
