package task

import (
	"context"

	"github.com/andremedeiros/loon/internal/project"
)

type DerivationCurrent struct{}

func (*DerivationCurrent) Header() string {
	return "Installing software"
}

func (*DerivationCurrent) Check(_ context.Context, p *project.Project) (bool, error) {
	return !p.NeedsUpdate(), nil
}

func (*DerivationCurrent) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	return p.EnsureDependencies()
}

func (*DerivationCurrent) Env(_ context.Context, _ *project.Project) (Env, BinPaths) {
	return nil, nil
}

func init() {
	RegisterTask("derivation:current:up", &DerivationCurrent{})
	RegisterTask("derivation:current:down", &DerivationCurrent{})
	RunsAfter("command:down", "derivation:current:down")
	RunsAfter("command:exec", "derivation:current:up")
	RunsAfter("command:task", "derivation:current:up")
	RunsAfter("command:up", "derivation:current:up")
}
