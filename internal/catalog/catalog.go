package catalog

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

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

func EntryFor(software string, version string) (Entry, error) {
	for _, e := range List() {
		if e.Name == software && e.Version == version {
			return e, nil
		}
	}

	return Entry{}, fmt.Errorf("no catalog entry for %s/%s", software, version)
}

func List() []Entry {
	es := []Entry{}

	assets := AssetNames()
	sort.Strings(assets)

	for _, asset := range assets {
		b, _ := Asset(asset)

		parts := strings.SplitN(asset, "/", 2)
		ext := filepath.Ext(parts[1])
		e := Entry{
			Name:    parts[0],
			Version: parts[1][0 : len(parts[1])-len(ext)],
		}
		json.Unmarshal(b, &e.Packages)
		es = append(es, e)
	}

	return es
}

func (pl PackageList) UnmarshalJSON(b []byte) error {
	pkgs := map[string]map[string]string{}
	json.Unmarshal(b, &pkgs)

	for name, pkg := range pkgs {
		pl = append(pl, Package{
			Package: name,
			Version: pkg["version"],
			URL:     pkg["url"],
			SHA256:  pkg["sha256"],
		})
	}

	return nil
}
