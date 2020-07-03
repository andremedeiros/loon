package task

import "github.com/andremedeiros/loon/internal/project"

type DerivationCurrent struct{}

func (*DerivationCurrent) Header() string {
	return "Installing software"
}

func (*DerivationCurrent) Check(p *project.Project) (bool, error) {
	return !p.NeedsUpdate(), nil
}

func (*DerivationCurrent) Resolve(p *project.Project) error {
	return p.EnsureDependencies()
}

func init() {
	RegisterTask("derivation:current", &DerivationCurrent{})
	Depends("derivation:current", "command:up")
	Depends("derivation:current", "command:exec")
	Depends("derivation:current", "command:task")
}
