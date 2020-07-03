package task

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andremedeiros/loon/internal/project"
)

func hbaConf(p *project.Project) []byte {
	buf := bytes.Buffer{}
	fmt.Fprintf(&buf, "host\tall\tall\t%s/32\ttrust", p.IP)
	return buf.Bytes()
}

type PostgresInitialize struct{}

func (*PostgresInitialize) Header() string {
	return "Initializing {blue:Postgres}"
}

func (*PostgresInitialize) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Postgres") {
		return true, nil
	}
	data := p.VariableDataPath("data", "postgres", "PG_VERSION")
	_, err := os.Stat(data)
	return err == nil, nil
}

func (*PostgresInitialize) Resolve(p *project.Project) error {
	data := p.VariableDataPath("data", "postgres")
	exe := p.DerivationExecutor()
	if err := exe.Execute([]string{"initdb", data}); err != nil {
		return err
	}
	hba := p.VariableDataPath("data", "postgres", "pg_hba.conf")
	return ioutil.WriteFile(hba, hbaConf(p), 0600)
}

type PostgresStart struct{}

func (*PostgresStart) Header() string {
	return "Starting {blue:Postgres}"
}

func (*PostgresStart) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Postgres") {
		return true, nil
	}
	return checkHealth(p.IP, 5432), nil
}

func (*PostgresStart) Resolve(p *project.Project) error {
	data := p.VariableDataPath("data", "postgres")
	log := p.VariableDataPath("data", "postgres", "postgres.log")
	socket := p.VariableDataPath("sockets")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("-o '-h %s'", p.IP),
		fmt.Sprintf("-o '--unix_socket_directories=%s'", socket),
		fmt.Sprintf("--pgdata=%s", data),
		fmt.Sprintf("--log=%s", log),
		"start",
	})
}

type PostgresStop struct{}

func (*PostgresStop) Header() string {
	return "Stopping {blue:Postgres}"
}

func (*PostgresStop) Check(p *project.Project) (bool, error) {
	if !checkProjectHasService(p, "Postgres") {
		return true, nil
	}
	return !checkHealth(p.IP, 5432), nil
}

func (*PostgresStop) Resolve(p *project.Project) error {
	data := p.VariableDataPath("data", "postgres")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("--pgdata=%s", data),
		"stop",
	})
}

func init() {
	RegisterTask("postgres:initialize", &PostgresInitialize{})
	RegisterTask("postgres:start", &PostgresStart{})
	RegisterTask("postgres:stop", &PostgresStop{})
	Depends("postgres:initialize", "derivation:current")
	Depends("postgres:start", "postgres:initialize")
	Depends("postgres:start", "networking:start")
	Depends("postgres:stop", "derivation:current")
	Depends("postgres:stop", "command:down")
}
