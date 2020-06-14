package catalog

type Service interface {
	String() string

	Install() error
	Uninstall() error

	Environ(string, string) []string

	HealthCheck() error
	Start(string, string) []string
	Stop(string, string) error
}

var Services = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
