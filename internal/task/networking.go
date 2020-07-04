package task

import (
	"context"
	"fmt"
	"net"

	"github.com/andremedeiros/loon/internal/project"
)

type Networking struct{}

func (*Networking) Header() string {
	return "Setting up project networking"
}

func (n *Networking) Check(_ context.Context, p *project.Project) (bool, error) {
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

func (*Networking) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	return []string{fmt.Sprintf("PROJECT_IP=%s", p.IP)}, nil
}

func init() {
	RegisterTask("networking:start", &Networking{})
	RunsAfter("command:up", "networking:start")
	RunsAfter("command:task", "networking:start")
	RunsAfter("command:exec", "networking:start")
}
