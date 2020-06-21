package service

type Executer interface {
	Execute([]string) error
}

type Service interface {
	String() string
	Identifier() string
	Environ(string, string) []string
	Initialize(Executer, string, string) error
	Start(Executer, string, string) error
	Stop(Executer, string, string) error
}

var Services = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
