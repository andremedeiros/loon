package task

import (
	"bytes"
	"context"
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

func (*PostgresInitialize) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "postgresql") {
		return true, nil
	}
	data := p.VariableDataPath("data", "postgres", "PG_VERSION")
	_, err := os.Stat(data)
	return err == nil, nil
}

func (*PostgresInitialize) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	data := p.VariableDataPath("data", "postgres")
	exe := p.DerivationExecutor()
	if err := exe.Execute([]string{"initdb", data}); err != nil {
		return err
	}
	hba := p.VariableDataPath("data", "postgres", "pg_hba.conf")
	return ioutil.WriteFile(hba, hbaConf(p), 0600)
}

func (*PostgresInitialize) Env(_ context.Context, _ *project.Project) (Env, BinPaths) {
	return nil, nil
}

type PostgresStart struct {
	started bool
}

func (*PostgresStart) Header() string {
	return "Starting {blue:Postgres}"
}

func (ps *PostgresStart) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "postgresql") {
		return true, nil
	}
	return checkHealth(p.IP, 5432, ps.started), nil
}

func (ps *PostgresStart) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	ps.started = true
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

func (*PostgresStart) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
	if checkProjectHasDep(p, "postgresql") {
		return Env{"DATABASE_URL": fmt.Sprintf("postgres://%s:5432", p.IP)}, nil
	}
	return nil, nil
}

type PostgresStop struct{}

func (*PostgresStop) Header() string {
	return "Stopping {blue:Postgres}"
}

func (*PostgresStop) Check(_ context.Context, p *project.Project) (bool, error) {
	if !checkProjectHasDep(p, "postgresql") {
		return true, nil
	}
	return checkDown(p.IP, 5432, false), nil
}

func (*PostgresStop) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	data := p.VariableDataPath("data", "postgres")
	exe := p.DerivationExecutor()
	return exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("--pgdata=%s", data),
		"stop",
	})
}

func (*PostgresStop) Env(_ context.Context, _ *project.Project) (Env, BinPaths) {
	return nil, nil
}

func init() {
	RegisterTask("postgres:initialize", &PostgresInitialize{})
	RegisterTask("postgres:start", &PostgresStart{})
	RegisterTask("postgres:stop", &PostgresStop{})
	RunsAfter("command:down", "postgres:stop")
	RunsAfter("derivation:current:down", "postgres:stop")
	RunsAfter("derivation:current:up", "postgres:initialize")
	RunsAfter("networking:start", "postgres:start")
	RunsAfter("postgres:initialize", "postgres:start")
}
