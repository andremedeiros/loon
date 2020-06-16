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

func (r *Redis) Initialize(_, _ string) []string {
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

func (r *Redis) Start(ipaddr, vdpath string) []string {
	pidsPath := filepath.Join(vdpath, "pids")
	os.MkdirAll(pidsPath, 0755)
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")

	vdPath := filepath.Join(vdpath, "redis")
	os.MkdirAll(vdPath, 0755)

	return []string{
		"redis-server",
		"--daemonize yes",
		"--port 6379",
		fmt.Sprintf("--dir %s", vdPath),
		fmt.Sprintf("--pidfile %s", pidPath),
	}
}

func (r *Redis) Stop(ipaddr, vdpath string) error {
	pidPath := filepath.Join(vdpath, "pids", "redis.pid")
	p, _ := process.FromPidFile(pidPath)
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
