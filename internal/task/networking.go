package task

import (
	"context"
	"net"

	"github.com/andremedeiros/loon/internal/project"
)

func checkIp(p *project.Project) (bool, error) {
	ifis, err := net.Interfaces()
	if err != nil {
		return false, err
	}
	for _, ifi := range ifis {
		addrs, err := ifi.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ipnet.IP.Equal(p.IP) {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

type NetworkingStart struct{}

func (*NetworkingStart) Header() string {
	return "Setting up project networking"
}

func (*NetworkingStart) Check(_ context.Context, p *project.Project) (bool, error) {
	return checkIp(p)
}

func (*NetworkingStart) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
	return Env{"PROJECT_IP": p.IP.String()}, nil
}

type NetworkingStop struct{}

func (*NetworkingStop) Header() string {
	return "Tearing down project networking"
}

func (*NetworkingStop) Check(_ context.Context, p *project.Project) (bool, error) {
	exists, err := checkIp(p)
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
