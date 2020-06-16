package catalog

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/process"
)

type Mysql struct{}

func (m *Mysql) String() string {
	return "MySQL"
}

func (m *Mysql) Initialize(ipaddr, vdpath string) []string {
	vdPath := filepath.Join(vdpath, "mysql")
	if _, err := os.Stat(vdPath); err == nil {
		return nil
	}
	os.MkdirAll(vdPath, 0755)
	return []string{
		"mysqld",
		"--initialize-insecure",
		fmt.Sprintf("--datadir=%s", vdPath),
	}
}

func (m *Mysql) Versions() map[string]Entry {
	return map[string]Entry{
		"default": EntryFor("mysql", "8.0.17"),
		"latest":  EntryFor("mysql", "8.0.17"),

		"8.0.17": EntryFor("mysql", "8.0.17"),
	}
}

func (m *Mysql) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("DATABASE_URL=%s:3306", ipaddr),
	}
}

func (m *Mysql) Start(ipaddr, vdpath string) []string {
	pidsPath := filepath.Join(vdpath, "pids")
	os.MkdirAll(pidsPath, 0755)
	pidPath := filepath.Join(vdpath, "pids", "mysql.pid")
	socketPath := filepath.Join(vdpath, "mysqld.sock")

	vdPath := filepath.Join(vdpath, "mysql")

	return []string{
		"mysqld",
		"--daemonize",
		fmt.Sprintf("--pid-file=%s", pidPath),
		fmt.Sprintf("--datadir=%s", vdPath),
		fmt.Sprintf("--bind-address=%s", ipaddr),
		fmt.Sprintf("--socket=%s", socketPath),
	}
}

func (m *Mysql) Stop(ipaddr, vdpath string) error {
	// TODO

	/*
			mysqladmin
			-u root
			--socket=/Users/andremedeiros/src/github.com/andremedeiros
		/loon/.loon/mysqld.sock shutdown
	*/
	pidPath := filepath.Join(vdpath, "pids", "mysql.pid")
	p, _ := process.FromPidFile(pidPath)
	_ = os.Remove(pidPath)
	return p.Signal(os.Interrupt)
}
