package provider

import "github.com/andremedeiros/loon/catalog"

type Provider interface {
	String() string
	Add(catalog.Entry) error
	Install() error
	Start(cmd []string) error
}

var Handlers = map[string]Provider{
	"nix": NewNix(),
}
