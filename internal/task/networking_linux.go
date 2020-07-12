package task

import (
	"context"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func (*NetworkingStart) Resolve(_ context.Context, p *project.Project) error {
	return executor.Execute([]string{
		"sudo",
		"ip",
		"addr",
		"add",
		p.IP.String(),
		"dev",
		"lo",
	})
}

func (*NetworkingStop) Resolve(_ context.Context, p *project.Project) error {
	return executor.Execute([]string{
		"sudo",
		"ip",
		"addr",
		"del",
		p.IP.String(),
		"dev",
		"lo",
	})
}
