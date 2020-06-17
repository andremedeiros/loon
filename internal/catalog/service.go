package catalog

import (
	"github.com/andremedeiros/loon/internal/nix"
)

type Installable interface {
	Versions() map[string]Entry
}

type Service interface {
	String() string
	Identifier() string
	Environ(string, string) []string
	Initialize(Executer, string, string) error
	Start(Executer, string, string) error
	Stop(Executer, string, string) error
}

func Packages(i Installable, version string) []nix.Package {
	pkgs := []nix.Package{}

	versions := i.Versions()
	entry, ok := versions[version]
	if !ok {
		return pkgs
	}

	for _, pkg := range entry.Packages {
		nixpkg := nix.Package{
			Name:    pkg.Package,
			Version: pkg.Version,
			URL:     pkg.URL,
			SHA256:  pkg.SHA256,
		}
		pkgs = append(pkgs, nixpkg)
	}

	return pkgs
}

var Services = map[string]Service{
	"memcached":  &Memcached{},
	"mysql":      &Mysql{},
	"postgresql": &Postgres{},
	"redis":      &Redis{},
}
