package task

import (
	"fmt"

	"github.com/andremedeiros/loon/internal/process"
	"github.com/andremedeiros/loon/internal/project"
)

type MemcachedStart struct{}

func (*MemcachedStart) Header() string {
	return "Starting {blue:Memcached}"
}

func (*MemcachedStart) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Memcached") {
		return true, nil
	}
	return checkHealth(p.IP, 11211), nil
}

func (*MemcachedStart) Resolve(p *project.Project) error {
	pid := p.VariableDataPath("pids", "memcached.pid")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"memcached",
		"--daemon",
		"--port=11211",
		fmt.Sprintf("--listen=%s", p.IP),
		fmt.Sprintf("--pidfile=%s", pid),
	})
}

type MemcachedStop struct {
	killed bool
}

func (*MemcachedStop) Header() string {
	return "Stopping {blue:Memcached}"
}

func (ms *MemcachedStop) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Memcached") || ms.killed {
		return true, nil
	}
	return !checkHealth(p.IP, 11211), nil
}

func (ms *MemcachedStop) Resolve(p *project.Project) error {
	ms.killed = true
	pid := p.VariableDataPath("pids", "memcached.pid")
	return process.InterruptFromPidFile(pid)
}

func init() {
	RegisterTask("memcached:start", &MemcachedStart{})
	RegisterTask("memcached:stop", &MemcachedStop{})
	Depends("memcached:start", "networking:start")
	Depends("memcached:stop", "command:down")
	Depends("memcached:stop", "derivation:current")
}
