package service

type Service interface {
	String() string
	Environ() []string
	Start() []string
	Stop() error
}

var Handlers = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
