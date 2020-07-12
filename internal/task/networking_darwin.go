package task

import (
	"context"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func (*NetworkingStart) Resolve(_ context.Context, p *project.Project) error {
	return executor.Execute([]string{
		"sudo",
		"ifconfig",
		"lo0",
		"alias",
		p.IP.String(),
		"255.255.255.0",
	})
}

func (*NetworkingStop) Resolve(_ context.Context, p *project.Project) error {
	return executor.Execute([]string{
		"sudo",
		"ifconfig",
		"lo0",
		"-alias",
		p.IP.String(),
	})
}
