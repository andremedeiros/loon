package task

import (
	"context"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func npmEnv(p *project.Project) Env {
	return Env{"NPM_CONFIG_PREFIX": p.VariableDataPath("data", "npm")}
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

func (*NodeInitialize) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	exe := p.DerivationExecutor()
	return exe.Execute(
		[]string{"npm", "install", "-g", "yarn"},
		executor.WithEnv(npmEnv(p)),
	)
}

func (*NodeInitialize) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
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
