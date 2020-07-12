package task

import (
	"context"
	"fmt"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func (*NetworkingStart) Resolve(_ context.Context, p *project.Project, sf SudoFunc) error {
	msg := `I need superuser access in order to {bold,underline:add an IP alias for the current project}.`
	sudo := sf(msg)
	return executor.Execute([]string{
		"ip",
		"addr",
		"add",
		fmt.Sprintf("%s/32", p.IP),
		"dev",
		"lo",
	}, executor.WithSudo(sudo))
}

func (*NetworkingStop) Resolve(_ context.Context, p *project.Project, sf SudoFunc) error {
	msg := `I need superuser access in order to {bold,underline:remove the IP alias for the current project}.`
	sudo := sf(msg)
	return executor.Execute([]string{
		"ip",
		"addr",
		"del",
		fmt.Sprintf("%s/32", p.IP),
		"dev",
		"lo",
	}, executor.WithSudo(sudo))
}
