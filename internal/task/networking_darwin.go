package task

import (
	"context"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func (*NetworkingStart) Resolve(_ context.Context, p *project.Project, sudo SudoFunc) error {
	msg := `I need superuser access in order to {bold:add an IP alias for the current project}.`

	return executor.Execute([]string{
		"ifconfig",
		"lo0",
		"alias",
		p.IP.String(),
		"255.255.255.0",
	}, executor.WithSudo(sudo(msg)))
}

func (*NetworkingStop) Resolve(_ context.Context, p *project.Project, sudo SudoFunc) error {
	msg := `I need superuser access in order to {bold:add an IP alias for the current project}.`

	return executor.Execute([]string{
		"ifconfig",
		"lo0",
		"-alias",
		p.IP.String(),
	}, executor.WithSudo(sudo(msg)))
}
