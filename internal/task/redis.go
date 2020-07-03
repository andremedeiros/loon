package task

import (
	"fmt"
	"os"

	"github.com/andremedeiros/loon/internal/process"
	"github.com/andremedeiros/loon/internal/project"
)

type RedisInitialize struct{}

func (*RedisInitialize) Header() string {
	return "Initializing {blue:Redis}"
}

func (*RedisInitialize) Check(p *project.Project) (bool, error) {
	data := p.VariableDataPath("data", "redis")
	_, err := os.Stat(data)
	return err == nil, nil
}

func (*RedisInitialize) Resolve(p *project.Project) error {
	data := p.VariableDataPath("data", "redis")
	return os.MkdirAll(data, 0755)
}

type RedisStart struct{}

func (*RedisStart) Header() string {
	return "Starting {blue:Redis}"
}

func (*RedisStart) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Redis") {
		return true, nil
	}
	return checkHealth(p.IP, 6379), nil
}

func (*RedisStart) Resolve(p *project.Project) error {
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

type RedisStop struct {
	killed bool
}

func (*RedisStop) Header() string {
	return "Stopping {blue:Redis}"
}

func (rs *RedisStop) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Redis") || rs.killed {
		return true, nil
	}
	return !checkHealth(p.IP, 6379), nil
}

func (rs *RedisStop) Resolve(p *project.Project) error {
	rs.killed = true
	pid := p.VariableDataPath("pids", "redis.pid")
	return process.InterruptFromPidFile(pid)
}

func init() {
	RegisterTask("redis:initialize", &RedisInitialize{})
	RegisterTask("redis:start", &RedisStart{})
	RegisterTask("redis:stop", &RedisStop{})
	Depends("redis:initialize", "derivation:current")
	Depends("redis:start", "redis:initialize")
	Depends("redis:start", "networking:start")
	Depends("redis:stop", "derivation:current")
	Depends("redis:stop", "command:down")
}
