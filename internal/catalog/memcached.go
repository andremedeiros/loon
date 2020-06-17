package catalog

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/process"
)

type Memcached struct{}

func (m *Memcached) String() string {
	return "Memcached"
}

func (m *Memcached) Identifier() string {
	return "memcached"
}

func (m *Memcached) Initialize(_ Executer, _, _ string) error {
	return nil
}

func (m *Memcached) Versions() map[string]Entry {
	return map[string]Entry{
		"default": EntryFor("memcached", "1.6.6"),
		"latest":  EntryFor("memcached", "1.6.6"),

		"1.6.6": EntryFor("memcached", "1.6.6"),
		"1.6.5": EntryFor("memcached", "1.6.5"),
	}
}

func (m *Memcached) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("MEMCACHED_URL=%s:11211", ipaddr),
	}
}

func (m *Memcached) Start(exe Executer, ipaddr, vdpath string) error {
	pidPath := filepath.Join(vdpath, "pids", "memcached.pid")

	return exe.Execute([]string{
		"memcached",
		"--daemon",
		"--port=11211",
		fmt.Sprintf("--listen=%s", ipaddr),
		fmt.Sprintf("--pidfile=%s", pidPath),
	})
}

func (m *Memcached) Stop(exe Executer, ipaddr, vdpath string) error {
	pidPath := filepath.Join(vdpath, "pids", "memcached.pid")
	p, err := process.FromPidFile(pidPath)
	if err != nil {
		return nil
	}
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
