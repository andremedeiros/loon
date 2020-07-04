package project

import "github.com/andremedeiros/loon/internal/nix"

type Dependency struct {
	Name     string
	Version  string
	Packages []nix.Package
}
