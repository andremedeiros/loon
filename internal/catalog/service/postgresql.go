package service

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/andremedeiros/loon/internal/executer"
)

type Postgres struct{}

func (p *Postgres) String() string {
	return "Postgres"
}

func (p *Postgres) Identifier() string {
	return "postgres"
}

func (p *Postgres) Initialize(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	dataPath := filepath.Join(vdpath, "data", "postgres")
	if _, err := os.Stat(filepath.Join(dataPath, "PG_VERSION")); err == nil {
		return nil
	}
	if _, err := exe.Execute([]string{"initdb", dataPath}, opts...); err != nil {
		return err
	}
	hbaConf := fmt.Sprintf("host\tall\tall\t%s/32\ttrust", ipaddr)
	return ioutil.WriteFile(filepath.Join(dataPath, "pg_hba.conf"), []byte(hbaConf), 0600)
}

func (p *Postgres) Versions() map[string][]string {
	return map[string][]string{
		"default": {"postgresql", "12.3"},
		"latest":  {"postgresql", "12.3"},

		"12.3":   {"postgresql", "12.3"},
		"11.8":   {"postgresql", "11.8"},
		"10.13":  {"postgresql", "10.13"},
		"9.6.18": {"postgresql", "9.6.18"},
		"9.5.22": {"postgresql", "9.5.22"},
	}
}

func (p *Postgres) Environ(ipaddr, vdpath string) []string {
	return []string{
		fmt.Sprintf("DATABASE_URL=postgres://%s:5432", ipaddr),
	}
}

func (p *Postgres) IsHealthy(ipaddr, _ string) bool {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:5432", ipaddr), 100*time.Millisecond)
	return err == nil
}

func (p *Postgres) Start(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	dataPath := filepath.Join(vdpath, "data", "postgres")
	logFilePath := filepath.Join(vdpath, "data", "postgres", "postgres.log")
	socketPath := filepath.Join(vdpath, "sockets")
	_, err := exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("-o '-h %s'", ipaddr),
		fmt.Sprintf("-o '--unix_socket_directories=%s'", socketPath),
		fmt.Sprintf("--pgdata=%s", dataPath),
		fmt.Sprintf("--log=%s", logFilePath),
		"start",
	}, opts...)
	return err
}

func (p *Postgres) Stop(exe executer.Executer, ipaddr, vdpath string, opts ...executer.Option) error {
	dataPath := filepath.Join(vdpath, "data", "postgres")
	_, err := exe.Execute([]string{
		"pg_ctl",
		fmt.Sprintf("--pgdata=%s", dataPath),
		"stop",
	}, opts...)
	return err
}
