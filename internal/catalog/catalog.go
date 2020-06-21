package catalog

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/andremedeiros/loon/internal/catalog/language"
	"github.com/andremedeiros/loon/internal/catalog/service"
	"github.com/andremedeiros/loon/internal/nix"
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

	assets := map[string][]byte{}
	for _, an := range service.AssetNames() {
		assets[an], _ = service.Asset(an)
	}
	for _, an := range language.AssetNames() {
		assets[an], _ = language.Asset(an)
	}

	for asset, b := range assets {
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

type Installable interface {
	Versions() map[string][]string
}

func Packages(i Installable, version string) []nix.Package {
	pkgs := []nix.Package{}

	versions := i.Versions()
	parts, ok := versions[version]
	if !ok {
		return pkgs
	}

	entry := EntryFor(parts[0], parts[1])
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
