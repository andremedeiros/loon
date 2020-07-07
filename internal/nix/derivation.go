package nix

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/andremedeiros/loon/internal/executor"
)

type Derivation struct {
	Packages []Package
	NixPath  string
	DrvPath  string
}

func NewDerivation(vdpath string) *Derivation {
	nixPath := filepath.Join(vdpath, "default.nix")
	drvPath := filepath.Join(vdpath, "loon.drv")
	return &Derivation{NixPath: nixPath, DrvPath: drvPath}
}

func (d *Derivation) Execute(args []string, opts ...executor.Option) error {
	cmd := []string{
		"nix-shell",
		d.DrvPath,
		"--command",
		strings.Join(args, " "),
	}
	return executor.Execute(cmd, opts...)
}

func (d *Derivation) generate() {
	tmpl := `
{ pkgs ? import <nixpkgs> { } }:
let
	inherit (pkgs) stdenv fetchurl mkShell;

	{{ range $package := .Packages }}
		{{ $package.Derivation }}
	{{ end }}
in mkShell {
	name = "loon";
	buildInputs = [
		{{ range $package := .Packages }}
			{{ $package.DerivationPackages }}
		{{ end }}
	];
}`
	buf := bytes.NewBuffer([]byte{})
	t, err := template.New("nix").Parse(tmpl)
	if err != nil {
		// TODO(andremedeiros): figure out a better way
		panic(err)
	}
	t.Execute(buf, d)
	fd, _ := os.OpenFile(d.NixPath, os.O_RDWR|os.O_CREATE, 0660)
	defer fd.Close()
	fd.Truncate(0)
	fd.Write(buf.Bytes())
}

func (d *Derivation) NeedsUpdate(payloadModified time.Time) bool {
	nixModified, err := os.Stat(d.NixPath)
	if err != nil {
		return true
	}
	return payloadModified.After(nixModified.ModTime())
}

func (d *Derivation) Install() error {
	d.generate()
	if err := executor.Execute([]string{"nix-instantiate", d.NixPath, "--indirect", "--add-root", d.DrvPath}); err != nil {
		return err
	}
	return d.Execute([]string{"true"}, executor.WithStdout(os.Stdout))
}

func (d *Derivation) Path() string {
	return d.NixPath
}
