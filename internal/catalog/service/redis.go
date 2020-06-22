package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/executer"
	"github.com/andremedeiros/loon/internal/process"
)

type Redis struct{}

func (r *Redis) String() string {
	return "Redis"
}

func (r *Redis) Identifier() string {
	return "redis"
}

func (r *Redis) Initialize(_ executer.Executer, _, _ string, _ ...executer.Option) error {
	return nil
}

func (r *Redis) Versions() map[string][]string {
	return map[string][]string{
		"default": {"redis", "6.0.4"},
		"latest":  {"redis", "6.0.4"},

		"6.0.4": {"redis", "6.0.4"},
	}
}

func (r *Redis) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("REDIS_URL=redis://%s:6379", ipaddr),
	}
}

func (r *Redis) Start(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	dataPath := filepath.Join(vdpath, "data", "redis")

	_, err := exe.Execute([]string{
		"redis-server",
		"--daemonize yes",
		"--port 6379",
		fmt.Sprintf("--bind %s", ipaddr),
		fmt.Sprintf("--dir %s", dataPath),
		fmt.Sprintf("--pidfile %s", pidPath),
	}, opts...)
	return err
}

func (r *Redis) Stop(exe executer.Executer, ipaddr, vdpath string, _ ...executer.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	p, err := process.FromPidFile(pidPath)
	if err != nil {
		return nil
	}
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
