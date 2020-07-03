package task

import (
	"fmt"
	"os"

	"github.com/andremedeiros/loon/internal/project"
)

type MysqlInitialize struct{}

func (*MysqlInitialize) Header() string {
	return "Initializing {blue:Mysql}"
}

func (*MysqlInitialize) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Mysql") {
		return true, nil
	}
	data := p.VariableDataPath("data", "mysql", "auto.cnf")
	_, err := os.Stat(data)
	return err == nil, nil
}
func (*MysqlInitialize) Resolve(p *project.Project) error {
	data := p.VariableDataPath("data", "mysql")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"mysqld",
		"--initialize-insecure",
		fmt.Sprintf("--datadir=%s", data),
	})
}

type MysqlStart struct{}

func (*MysqlStart) Header() string {
	return "Starting {blue:Mysql}"
}

func (*MysqlStart) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Mysql") {
		return true, nil
	}
	return checkHealth(p.IP, 3306), nil
}

func (*MysqlStart) Resolve(p *project.Project) error {
	pid := p.VariableDataPath("pids", "mysql.pid")
	data := p.VariableDataPath("data", "mysql")
	socket := p.VariableDataPath("sockets", "mysqld.sock")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"mysqld",
		"--daemonize",
		fmt.Sprintf("--pid-file=%s", pid),
		fmt.Sprintf("--datadir=%s", data),
		fmt.Sprintf("--bind-address=%s", p.IP),
		fmt.Sprintf("--socket=%s", socket),
	})
}

type MysqlStop struct{}

func (*MysqlStop) Header() string {
	return "Stopping {blue:Mysql}"
}

func (*MysqlStop) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Mysql") {
		return true, nil
	}
	return !checkHealth(p.IP, 3306), nil
}
func (*MysqlStop) Resolve(p *project.Project) error {
	socket := p.VariableDataPath("sockets", "mysqld.sock")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"mysqladmin",
		"-u root",
		fmt.Sprintf("--socket=%s", socket),
		"shutdown",
	})
}

func init() {
	RegisterTask("mysql:initialize", &MysqlInitialize{})
	RegisterTask("mysql:start", &MysqlStart{})
	RegisterTask("mysql:stop", &MysqlStop{})
	Depends("mysql:initialize", "derivation:current")
	Depends("mysql:start", "mysql:initialize")
	Depends("mysql:start", "networking:start")
	Depends("mysql:stop", "derivation:current")
	Depends("mysql:stop", "command:down")
}
