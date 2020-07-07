package task

import (
	"context"
	"fmt"
	"os"

	"github.com/andremedeiros/loon/internal/process"
	"github.com/andremedeiros/loon/internal/project"
)

type RedisInitialize struct{}

func (*RedisInitialize) Header() string {
	return "Initializing {blue:Redis}"
}

func (*RedisInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "redis") {
		return true, nil
	}
	data := p.VariableDataPath("data", "redis")
	_, err := os.Stat(data)
	return err == nil, nil
}

func (*RedisInitialize) Resolve(_ context.Context, p *project.Project) error {
	data := p.VariableDataPath("data", "redis")
	return os.MkdirAll(data, 0755)
}

func (*RedisInitialize) Environ(_ context.Context, _ *project.Project) (Environ, BinPaths) {
	return nil, nil
}

type RedisStart struct {
	started bool
}

func (*RedisStart) Header() string {
	return "Starting {blue:Redis}"
}

func (rs *RedisStart) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "redis") {
		return true, nil
	}
	return checkHealth(p.IP, 6379, rs.started), nil
}

func (rs *RedisStart) Resolve(_ context.Context, p *project.Project) error {
	rs.started = true
	pid := p.VariableDataPath("pids", "redis.pid")
	data := p.VariableDataPath("data", "redis")
	exe := p.DerivationExecutor()

	return exe.Execute([]string{
		"redis-server",
		"--daemonize yes",
		"--port 6379",
		fmt.Sprintf("--bind %s", p.IP.String()),
		fmt.Sprintf("--dir %s", data),
		fmt.Sprintf("--pidfile %s", pid),
	})
}

func (*RedisStart) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	if checkProjectHasDep(p, "redis") {
		return []string{fmt.Sprintf("REDIS_URL=redis://%s:6379", p.IP)}, nil
	}
	return nil, nil
}

type RedisStop struct {
	killed bool
}

func (*RedisStop) Header() string {
	return "Stopping {blue:Redis}"
}

func (rs *RedisStop) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "redis") {
		return true, nil
	}
	return checkDown(p.IP, 6379, rs.killed), nil
}

func (rs *RedisStop) Resolve(_ context.Context, p *project.Project) error {
	rs.killed = true
	pid := p.VariableDataPath("pids", "redis.pid")
	return process.InterruptFromPidFile(pid)
}

func (*RedisStop) Environ(_ context.Context, _ *project.Project) (Environ, BinPaths) {
	return nil, nil
}

func init() {
	RegisterTask("redis:initialize", &RedisInitialize{})
	RegisterTask("redis:start", &RedisStart{})
	RegisterTask("redis:stop", &RedisStop{})
	RunsAfter("command:down", "redis:stop")
	RunsAfter("derivation:current:down", "redis:stop")
	RunsAfter("derivation:current:up", "redis:initialize")
	RunsAfter("networking:start", "redis:start")
	RunsAfter("redis:initialize", "redis:start")
}
