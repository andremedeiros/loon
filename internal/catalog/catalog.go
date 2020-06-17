package catalog

import (
	"encoding/json"
	"path/filepath"
	"sort"
	"strings"
)

type Executer interface {
	Execute([]string) error
}

type Package struct {
	Package string
	Version string `json:"version"`
	URL     string `json:"url"`
	SHA256  string `json:"sha256"`
}

type PackageList []Package

type Entry struct {
	Name     string
	Version  string
	Packages PackageList
}

func EntryFor(software string, version string) Entry {
	for _, e := range List() {
		if e.Name == software && e.Version == version {
			return e
		}
	}

	panic("no entry")
}

func List() []Entry {
	es := []Entry{}

	assets := AssetNames()
	sort.Strings(assets)

	for _, asset := range assets {
		parts := strings.SplitN(asset, "/", 2)
		ext := filepath.Ext(parts[1])
		e := Entry{
			Name:    parts[0],
			Version: parts[1][0 : len(parts[1])-len(ext)],
		}

		{
			b, _ := Asset(asset)
			pkgs := map[string]map[string]string{}
			json.Unmarshal(b, &pkgs)
			for name, pkg := range pkgs {
				e.Packages = append(e.Packages, Package{
					Package: name,
					Version: pkg["version"],
					URL:     pkg["url"],
					SHA256:  pkg["sha256"],
				})
			}
		}

		es = append(es, e)
	}

	return es
}
