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

func (d *Derivation) Execute(args []string, opts ...executor.Option) (int, error) {
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
{{ range $package := $.Packages }}
	{{ $package.Nix }}
{{ end }}

in mkShell {
	buildInputs = [{{ range $package := .Packages }}
		{{ $package.Name }}
	{{ end }}];
}`

	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("nix").Parse(tmpl))
	t.Execute(buf, d)

	fd, _ := os.OpenFile(d.NixPath, os.O_RDWR|os.O_CREATE, 0660)
	defer fd.Close()
	fd.Truncate(0)
	fd.Write(buf.Bytes())
}

func (d *Derivation) NeedsUpdate(since time.Time) bool {
	fi, err := os.Stat(d.NixPath)
	if err != nil {
		return true
	}
	return fi.ModTime().Before(since)
}

func (d *Derivation) Install() error {
	d.generate()
	executor.Execute([]string{"nix-instantiate", d.NixPath, "--indirect", "--add-root", d.DrvPath})
	_, err := d.Execute([]string{"true"}, executor.WithStdout(os.Stdout))
	return err
}

func (d *Derivation) Path() string {
	return d.NixPath
}
