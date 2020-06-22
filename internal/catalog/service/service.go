package service

import "github.com/andremedeiros/loon/internal/executer"

type Service interface {
	String() string
	Identifier() string
	Environ(string, string) []string
	Initialize(executer.Executer, string, string, ...executer.Option) error
	Start(executer.Executer, string, string, ...executer.Option) error
	Stop(executer.Executer, string, string, ...executer.Option) error
}

var Services = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
