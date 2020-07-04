package task

import (
	"context"
	"fmt"

	"github.com/andremedeiros/loon/internal/process"
	"github.com/andremedeiros/loon/internal/project"
)

type MemcachedStart struct {
	started bool
}

func (*MemcachedStart) Header() string {
	return "Starting {blue:Memcached}"
}

func (ms *MemcachedStart) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "memcached") {
		return true, nil
	}
	return checkHealth(p.IP, 11211, ms.started), nil
}

func (ms *MemcachedStart) Resolve(_ context.Context, p *project.Project) error {
	ms.started = true
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

func (*MemcachedStart) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	if checkProjectHasDep(p, "memcached") {
		return []string{fmt.Sprintf("MEMCACHED_URL=%s:11211", p.IP)}, nil
	}
	return nil, nil
}

type MemcachedStop struct {
	killed bool
}

func (*MemcachedStop) Header() string {
	return "Stopping {blue:Memcached}"
}

func (ms *MemcachedStop) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "memcached") {
		return true, nil
	}
	return !checkHealth(p.IP, 11211, ms.killed), nil
}

func (ms *MemcachedStop) Resolve(_ context.Context, p *project.Project) error {
	ms.killed = true
	pid := p.VariableDataPath("pids", "memcached.pid")
	return process.InterruptFromPidFile(pid)
}

func (*MemcachedStop) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	return nil, nil
}

func init() {
	RegisterTask("memcached:start", &MemcachedStart{})
	RegisterTask("memcached:stop", &MemcachedStop{})
	RunsAfter("networking:start", "memcached:start")
	RunsAfter("command:down", "memcached:stop")
	RunsAfter("derivation:current:down", "memcached:stop")
}
