package service

import (
	"net"

	"github.com/andremedeiros/loon/internal/executor"
)

type Service interface {
	String() string
	Identifier() string
	Environ(net.IP, string) []string
	Initialize(executor.Executor, net.IP, string, ...executor.Option) error
	IsHealthy(net.IP, string) bool
	Start(executor.Executor, net.IP, string, ...executor.Option) error
	Stop(executor.Executor, net.IP, string, ...executor.Option) error
}

var Services = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
