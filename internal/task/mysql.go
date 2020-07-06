package task

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/andremedeiros/loon/internal/project"
)

const mysqlInit = `
	CREATE USER 'root'@'{{ .IP }}' IDENTIFIED BY '';
	GRANT ALL PRIVILEGES ON *.* TO 'root'@'{{ .IP }}';
	FLUSH PRIVILEGES;
`

func mysqlInitFile(p *project.Project) string {
	tmpfile, _ := ioutil.TempFile("", "init.sql")
	t := template.Must(template.New("init.sql").Parse(mysqlInit))
	t.Execute(tmpfile, p)
	tmpfile.Close()
	return tmpfile.Name()
}

type MysqlInitialize struct{}

func (*MysqlInitialize) Header() string {
	return "Initializing {blue:Mysql}"
}

func (*MysqlInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "mysql") {
		return true, nil
	}
	data := p.VariableDataPath("data", "mysql", "auto.cnf")
	_, err := os.Stat(data)
	return err == nil, nil
}
func (*MysqlInitialize) Resolve(_ context.Context, p *project.Project) error {
	data := p.VariableDataPath("data", "mysql")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"mysqld",
		"--initialize-insecure",
		fmt.Sprintf("--datadir=%s", data),
	})
}

func (*MysqlInitialize) Environ(_ context.Context, _ *project.Project) (Environ, BinPaths) {
	return nil, nil
}

type MysqlStart struct {
	started bool
}

func (*MysqlStart) Header() string {
	return "Starting {blue:Mysql}"
}

func (ms *MysqlStart) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "mysql") {
		return true, nil
	}
	return checkHealth(p.IP, 3306, ms.started), nil
}

func (ms *MysqlStart) Resolve(_ context.Context, p *project.Project) error {
	ms.started = true
	init := mysqlInitFile(p)
	defer os.Remove(init)
	pid := p.VariableDataPath("pids", "mysql.pid")
	data := p.VariableDataPath("data", "mysql")
	logerr := p.VariableDataPath("data", "mysql", "mysqld.err")
	socket := p.VariableDataPath("sockets", "mysqld.sock")
	xsocket := p.VariableDataPath("sockets", "mysqlx.sock")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"mysqld",
		"--daemonize",
		fmt.Sprintf("--pid-file=%s", pid),
		fmt.Sprintf("--datadir=%s", data),
		fmt.Sprintf("--bind-address=%s", p.IP),
		fmt.Sprintf("--mysqlx-bind-address=%s", p.IP),
		fmt.Sprintf("--socket=%s", socket),
		fmt.Sprintf("--mysqlx-socket=%s", xsocket),
		fmt.Sprintf("--init-file=%s", init),
		fmt.Sprintf("--log-error=%s", logerr),
	})
}

func (*MysqlStart) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	if checkProjectHasDep(p, "mysql") {
		return []string{fmt.Sprintf("DATABASE_URL=%s:3306", p.IP)}, nil
	}
	return nil, nil
}

type MysqlStop struct{}

func (*MysqlStop) Header() string {
	return "Stopping {blue:Mysql}"
}

func (*MysqlStop) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "mysql") {
		return true, nil
	}
	return checkDown(p.IP, 3306, false), nil
}

func (*MysqlStop) Resolve(_ context.Context, p *project.Project) error {
	socket := p.VariableDataPath("sockets", "mysqld.sock")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"mysqladmin",
		"-u root",
		fmt.Sprintf("--socket=%s", socket),
		"shutdown",
	})
}

func (*MysqlStop) Environ(_ context.Context, _ *project.Project) (Environ, BinPaths) {
	return nil, nil
}

func init() {
	RegisterTask("mysql:initialize", &MysqlInitialize{})
	RegisterTask("mysql:start", &MysqlStart{})
	RegisterTask("mysql:stop", &MysqlStop{})
	RunsAfter("command:down", "mysql:stop")
	RunsAfter("derivation:current:down", "mysql:stop")
	RunsAfter("derivation:current:up", "mysql:initialize")
	RunsAfter("mysql:initialize", "mysql:start")
	RunsAfter("networking:start", "mysql:start")
}
