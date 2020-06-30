package service

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/process"
)

type Redis struct{}

func (r *Redis) String() string {
	return "Redis"
}

func (r *Redis) Identifier() string {
	return "redis"
}

func (r *Redis) Initialize(_ executor.Executor, _ net.IP, _ string, _ ...executor.Option) error {
	return nil
}

func (r *Redis) Versions() map[string][]string {
	return map[string][]string{
		"default": {"redis", "6.0.4"},
		"latest":  {"redis", "6.0.4"},

		"6.0.4": {"redis", "6.0.4"},
	}
}

func (r *Redis) Environ(ip net.IP, _ string) []string {
	return []string{
		fmt.Sprintf("REDIS_URL=redis://%s:6379", ip),
	}
}

func (r *Redis) IsHealthy(ip net.IP, _ string) bool {
	hp := fmt.Sprintf("%s:6379", ip)
	_, err := net.DialTimeout("tcp", hp, 100*time.Millisecond)
	return err == nil
}

func (r *Redis) Start(exe executor.Executor, ip net.IP, vdpath string, opts ...executor.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	dataPath := filepath.Join(vdpath, "data", "redis")

	return exe.Execute([]string{
		"redis-server",
		"--daemonize yes",
		"--port 6379",
		fmt.Sprintf("--bind %s", ip),
		fmt.Sprintf("--dir %s", dataPath),
		fmt.Sprintf("--pidfile %s", pidPath),
	}, opts...)
}

func (r *Redis) Stop(exe executor.Executor, _ net.IP, vdpath string, _ ...executor.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	p, err := process.FromPidFile(pidPath)
	if err != nil {
		return nil
	}
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
