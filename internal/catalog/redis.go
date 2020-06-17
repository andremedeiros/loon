package catalog

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/process"
)

type Redis struct{}

func (r *Redis) String() string {
	return "Redis"
}

func (r *Redis) Identifier() string {
	return "redis"
}

func (r *Redis) Initialize(_ Executer, _, _ string) error {
	return nil
}

func (r *Redis) Versions() map[string]Entry {
	return map[string]Entry{
		"default": EntryFor("redis", "6.0.4"),
		"latest":  EntryFor("redis", "6.0.4"),

		"6.0.4": EntryFor("redis", "6.0.4"),
	}
}

func (r *Redis) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("REDIS_URL=redis://%s:6379", ipaddr),
	}
}

func (r *Redis) Start(exe Executer, ipaddr, vdpath string) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	dataPath := filepath.Join(vdpath, "data", "redis")

	return exe.Execute([]string{
		"redis-server",
		"--daemonize yes",
		"--port 6379",
		fmt.Sprintf("--bind %s", ipaddr),
		fmt.Sprintf("--dir %s", dataPath),
		fmt.Sprintf("--pidfile %s", pidPath),
	})
}

func (r *Redis) Stop(exe Executer, ipaddr, vdpath string) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	p, err := process.FromPidFile(pidPath)
	if err != nil {
		return nil
	}
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
