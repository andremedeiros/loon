package task

import (
	"context"
	"os"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func rubyEnv(p *project.Project) Env {
	gems := p.VariableDataPath("data", "gem")
	return Env{
		"GEM_HOME": gems,
		"GEM_PATH": gems,
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

func (*RubyInitialize) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	exe := p.DerivationExecutor()
	cmds := [][]string{
		{"gem", "install", "bundler", "--no-document"},       // Installs bundler
		{"bundle", "config", "build.sassc", "--disable-lto"}, // https://github.com/sass/sassc-ruby/issues/148
	}
	for _, cmd := range cmds {
		if err := exe.Execute(cmd, executor.WithEnv(rubyEnv(p))); err != nil {
			return err
		}
	}
	return nil
}

func (*RubyInitialize) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
	if !checkProjectHasDep(p, "ruby") {
		return nil, nil
	}
	bin := p.VariableDataPath("data", "gem", "bin")
	return rubyEnv(p), []string{bin}
}

func init() {
	RegisterTask("ruby:initialize", &RubyInitialize{})
	RunsAfter("derivation:current:up", "ruby:initialize")
}
