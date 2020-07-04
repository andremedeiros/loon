package task

import (
	"context"
	"fmt"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func npmEnviron(p *project.Project) []string {
	npm := p.VariableDataPath("data", "npm")
	return []string{
		fmt.Sprintf("NPM_CONFIG_PREFIX=%s", npm),
	}
}

type NodeInitialize struct{}

func (*NodeInitialize) Header() string {
	return "Setting up {blue:Node.JS}"
}

func (*NodeInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "node") {
		return true, nil
	}
	yarn := p.VariableDataPath("data", "npm", "bin", "yarn")
	if _, err := os.Stat(yarn); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func (*NodeInitialize) Resolve(_ context.Context, p *project.Project) error {
	exe := p.DerivationExecutor()
	return exe.Execute(
		[]string{"npm", "install", "-g", "yarn"},
		executor.WithEnviron(npmEnviron(p)),
	)
}

func (*NodeInitialize) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	if !checkProjectHasDep(p, "node") {
		return nil, nil
	}
	bin := p.VariableDataPath("data", "npm", "bin")
	return nil, []string{bin}
}

func init() {
	RegisterTask("node:initialize", &NodeInitialize{})
	RunsAfter("derivation:current:up", "node:initialize")
}
