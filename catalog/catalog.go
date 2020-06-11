package catalog

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

type Entry struct {
	Provider string
	Software string
	Version  string

	Payload []byte
}

func EntryFor(prov string, software string, version string) (Entry, error) {
	for _, e := range List() {
		if e.Provider == prov && e.Software == software {
			return e, nil
		}
	}

	return Entry{}, fmt.Errorf("no catalog entry for %s/%s/%s", prov, software, version)
}

func List() []Entry {
	es := []Entry{}

	assets := AssetNames()
	sort.Strings(assets)

	for _, asset := range assets {
		b, _ := Asset(asset)
		parts := strings.SplitN(asset, "/", 3)
		ext := filepath.Ext(parts[2])
		ver := parts[2][0 : len(parts[2])-len(ext)]
		e := Entry{
			Provider: parts[0],
			Software: parts[1],
			Version:  ver,
			Payload:  b,
		}
		es = append(es, e)
	}

	return es
}
