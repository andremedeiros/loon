package service

import (
	"net"

	"github.com/andremedeiros/loon/internal/executer"
)

type Service interface {
	String() string
	Identifier() string
	Environ(net.IP, string) []string
	Initialize(executer.Executer, net.IP, string, ...executer.Option) error
	IsHealthy(net.IP, string) bool
	Start(executer.Executer, net.IP, string, ...executer.Option) error
	Stop(executer.Executer, net.IP, string, ...executer.Option) error
}

var Services = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
