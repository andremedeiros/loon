package catalog

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/andremedeiros/loon/internal/nix"
)

type PackageList []nix.Package

type Entry struct {
	Name     string
	Version  string
	Packages PackageList
}

func EntryFor(software string, version string) (Entry, error) {
	for _, e := range List() {
		if e.Name == software && e.Version == version {
			return e, nil
		}
	}

	return Entry{}, fmt.Errorf("%s %s not supported", software, version)
}

func List() []Entry {
	es := []Entry{}

	for _, asset := range AssetNames() {
		b := MustAsset(asset)
		parts := strings.SplitN(asset, "/", 2)
		ext := filepath.Ext(parts[1])
		e := Entry{
			Name:    parts[0],
			Version: parts[1][0 : len(parts[1])-len(ext)],
		}

		{
			pkgs := map[string]map[string]string{}
			json.Unmarshal(b, &pkgs)
			for name, pkg := range pkgs {
				p := nix.NewPackage(name, pkg)
				e.Packages = append(e.Packages, p)
			}
		}

		es = append(es, e)
	}

	return es
}

type Installable interface {
	Versions() map[string][]string
}

func Packages(name string, version string) ([]nix.Package, error) {
	e, err := EntryFor(name, version)
	return e.Packages, err
}
