package service

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/executer"
)

type Mysql struct{}

func (m *Mysql) String() string {
	return "MySQL"
}

func (m *Mysql) Identifier() string {
	return "mysql"
}

func (m *Mysql) Initialize(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	dataPath := filepath.Join(vdpath, "data", "mysql")
	if _, err := os.Stat(filepath.Join(dataPath, "auto.cnf")); err == nil {
		return nil
	}

	_, err := exe.Execute([]string{
		"mysqld",
		"--initialize-insecure",
		fmt.Sprintf("--datadir=%s", dataPath),
	}, opts...)
	return err
}

func (m *Mysql) Versions() map[string][]string {
	return map[string][]string{
		"default": {"mysql", "8.0.17"},
		"latest":  {"mysql", "8.0.17"},

		"8.0.17": {"mysql", "8.0.17"},
	}
}

func (m *Mysql) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("DATABASE_URL=%s:3306", ipaddr),
	}
}

func (m *Mysql) Start(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	pidPath := filepath.Join(vdpath, "pids", "mysql.pid")
	dataPath := filepath.Join(vdpath, "data", "mysql")
	socketPath := filepath.Join(vdpath, "sockets", "mysqld.sock")

	_, err := exe.Execute([]string{
		"mysqld",
		"--daemonize",
		fmt.Sprintf("--pid-file=%s", pidPath),
		fmt.Sprintf("--datadir=%s", dataPath),
		fmt.Sprintf("--bind-address=%s", ipaddr),
		fmt.Sprintf("--socket=%s", socketPath),
	}, opts...)
	return err
}

func (m *Mysql) Stop(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	socketPath := filepath.Join(vdpath, "sockets", "mysqld.sock")

	if _, err := os.Stat(socketPath); err != nil {
		return nil
	}

	_, err := exe.Execute([]string{
		"mysqladmin",
		"-u root",
		fmt.Sprintf("--socket=%s", socketPath),
		"shutdown",
	}, opts...)
	return err
}
