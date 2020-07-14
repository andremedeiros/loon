package task

import (
	"context"

	"github.com/andremedeiros/loon/internal/project"
)

type NetworkingStart struct{}

func (*NetworkingStart) Header() string {
	return "Setting up project networking"
}

func (*NetworkingStart) Check(_ context.Context, p *project.Project) (bool, error) {
	return checkIp(p.IP)
}

func (*NetworkingStart) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
	return Env{"LOON_PROJECT_IP": p.IP.String()}, nil
}

type NetworkingStop struct{}

func (*NetworkingStop) Header() string {
	return "Tearing down project networking"
}

func (*NetworkingStop) Check(_ context.Context, p *project.Project) (bool, error) {
	exists, err := checkIp(p.IP)
	return !exists, err
}

func (*NetworkingStop) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
	return nil, nil
}

func init() {
	RegisterTask("networking:start", &NetworkingStart{})
	RegisterTask("networking:stop", &NetworkingStop{})
	RunsAfter("command:up", "networking:start")
	RunsAfter("command:task", "networking:start")
	RunsAfter("command:exec", "networking:start")
	RunsAfter("command:down:done", "networking:stop")
}
