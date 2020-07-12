package task

import (
	"context"
	"fmt"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

func (*NetworkingStart) Resolve(_ context.Context, p *project.Project) error {
	return executor.Execute([]string{
		"sudo",
		"ip",
		"addr",
		"add",
		fmt.Sprintf("%s/32", p.IP),
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
		fmt.Sprintf("%s/32", p.IP),
		"dev",
		"lo",
	})
}
