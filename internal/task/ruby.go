package task

import (
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

func (*RubyInitialize) Check(p *project.Project) (bool, error) {
	if !checkProjectHasLanguage(p, "Ruby") {
		return true, nil
	}
	bundler := p.VariableDataPath("data", "gem", "bin", "bundle")
	if _, err := os.Stat(bundler); os.IsNotExist(err) {
		return false, nil
	}
	return true, nil
}

func (*RubyInitialize) Resolve(p *project.Project) error {
	exe := p.DerivationExecutor()
	return exe.Execute(
		[]string{"gem", "install", "bundler", "--no-document"},
		executor.WithEnviron(rubyEnviron(p)),
		executor.WithStdout(os.Stdout),
	)
}

func init() {
	RegisterTask("ruby:initialize", &RubyInitialize{})
	Depends("ruby:initialize", "derivation:current")
}
