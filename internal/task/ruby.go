package task

import (
	"context"
	"fmt"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func rubyEnviron(p *project.Project) []string {
	gems := p.VariableDataPath("data", "gem")
	return []string{
		fmt.Sprintf("GEM_HOME=%s", gems),
		fmt.Sprintf("GEM_PATH=%s", gems),
	}
}

type RubyInitialize struct{}

func (*RubyInitialize) Header() string {
	return "Setting up {blue:Ruby}"
}

func (*RubyInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "ruby") {
		return true, nil
	}
	bundler := p.VariableDataPath("data", "gem", "bin", "bundle")
	if _, err := os.Stat(bundler); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func (*RubyInitialize) Resolve(_ context.Context, p *project.Project) error {
	exe := p.DerivationExecutor()
	return exe.Execute(
		[]string{"gem", "install", "bundler", "--no-document"},
		executor.WithEnviron(rubyEnviron(p)),
	)
}

func (*RubyInitialize) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	if !checkProjectHasDep(p, "ruby") {
		return nil, nil
	}
	bin := p.VariableDataPath("data", "gem", "bin")
	return rubyEnviron(p), []string{bin}
}

func init() {
	RegisterTask("ruby:initialize", &RubyInitialize{})
	RunsAfter("derivation:current:up", "ruby:initialize")
}
