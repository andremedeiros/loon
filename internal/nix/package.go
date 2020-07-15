package nix

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/andremedeiros/loon/internal/catalog"
	"github.com/andremedeiros/loon/internal/executor"
)

type Package struct {
	Name     string
	Version  string
	Template string
}

func (p Package) Derivation() string {
	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("pkg").Parse(p.Template))
	t, _ = t.Parse(`{{ template "derivation" }}`)
	t.Execute(buf, nil)
	return buf.String()
}

func (p Package) DerivationPackages() string {
	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("pkg").Parse(p.Template))
	t, _ = t.Parse(`{{ template "packages" }}`)
	t.Execute(buf, nil)
	return buf.String()
}

func (p Package) Build() error {
	tmpl := `
{ pkgs ? import <nixpkgs> { } }:
let
	inherit (pkgs) stdenv fetchurl mkShell;
	{{ .Derivation }}
in rec {
	package =
		{{ .DerivationPackages }}
	;
}`
	buf := bytes.NewBuffer([]byte{})
	t, err := template.New("nix").Parse(tmpl)
	if err != nil {
		// TODO(andremedeiros): figure out a better way
		panic(err)
	}
	t.Execute(buf, p)
	f, _ := ioutil.TempFile("", "default.nix")
	f.Write(buf.Bytes())
	f.Close()

	return executor.Execute([]string{
		"nix-build",
		f.Name(),
		"-A",
		"package",
		"--no-out-link",
	},
		executor.WithStdout(os.Stdout),
		executor.WithStderr(os.Stderr),
	)
}

func PackageFor(name string, version string) (Package, error) {
	for _, p := range Packages() {
		if p.Name == name && p.Version == version {
			return p, nil
		}
	}
	return Package{}, fmt.Errorf("%s %s not supported", name, version)
}

func Packages() []Package {
	ps := []Package{}
	for _, asset := range catalog.AssetNames() {
		b, _ := catalog.Asset(asset)
		parts := strings.SplitN(asset, "/", 2)
		ext := filepath.Ext(parts[1])
		p := Package{
			Name:     parts[0],
			Version:  parts[1][0 : len(parts[1])-len(ext)],
			Template: string(b),
		}
		ps = append(ps, p)
	}
	return ps
}
