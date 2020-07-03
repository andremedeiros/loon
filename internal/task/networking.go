package task

import (
	"net"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/project"
)

type Networking struct{}

func (*Networking) Header() string {
	return "Setting up project networking"
}

func (n *Networking) Check(p *project.Project) (bool, error) {
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

func (n *Networking) Resolve(p *project.Project) error {
	return executor.Execute([]string{
		"sudo",
		"ifconfig",
		"lo0",
		"alias",
		p.IP.String(),
		"255.255.255.0",
	})
}

func init() {
	RegisterTask("networking:start", &Networking{})
	Depends("networking:start", "command:up")
	Depends("networking:start", "command:task")
	Depends("networking:start", "command:exec")
}
