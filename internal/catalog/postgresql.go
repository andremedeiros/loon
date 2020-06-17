package catalog

import (
	"fmt"
	"os"
	"path/filepath"
)

type Postgres struct{}

func (p *Postgres) String() string {
	return "Postgres"
}

func (p *Postgres) Identifier() string {
	return "postgres"
}

func (p *Postgres) Initialize(exe Executer, _, vdpath string) error {
	dataPath := filepath.Join(vdpath, "data", "postgres")
	if _, err := os.Stat(filepath.Join(dataPath, "PG_VERSION")); err == nil {
		return nil
	}

	return exe.Execute([]string{
		"initdb",
		dataPath,
	})
}

func (p *Postgres) Versions() map[string]Entry {
	return map[string]Entry{
		"default": EntryFor("postgresql", "12.3"),
		"latest":  EntryFor("postgresql", "12.3"),

		"12.3":   EntryFor("postgresql", "12.3"),
		"11.8":   EntryFor("postgresql", "11.8"),
		"10.13":  EntryFor("postgresql", "10.13"),
		"9.6.18": EntryFor("postgresql", "9.6.18"),
		"9.5.22": EntryFor("postgresql", "9.5.22"),
	}
}

func (p *Postgres) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("DATABASE_URL=%s:5432", ipaddr),
	}
}

func (p *Postgres) Start(exe Executer, ipaddr, vdpath string) error {
	dataPath := filepath.Join(vdpath, "data", "postgres")
	logFilePath := filepath.Join(vdpath, "data", "postgres", "postgres.log")
	socketPath := filepath.Join(vdpath, "sockets")

	return exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("-o '-h %s'", ipaddr),
		fmt.Sprintf("-o '--unix_socket_directories=%s'", socketPath),
		fmt.Sprintf("--pgdata=%s", dataPath),
		fmt.Sprintf("--log=%s", logFilePath),
		"start",
	})
}

func (p *Postgres) Stop(exe Executer, ipaddr, vdpath string) error {
	dataPath := filepath.Join(vdpath, "data", "postgres")
	return exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("--pgdata=%s", dataPath),
		"stop",
	})
}
