package task

import (
	"context"

	"github.com/andremedeiros/loon/internal/project"
)

func goEnviron(p *project.Project) map[string]string {
	return Env{"GOPATH": p.VariableDataPath("data", "go")}
}

type GoInitialize struct{}

func (*GoInitialize) Header() string {
	return "Setting up {blue:Golang}"
}

func (*GoInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	return true, nil
}

func (*GoInitialize) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	return nil
}

func (*GoInitialize) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
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
