package nix

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/andremedeiros/loon/internal/executer"
)

type Derivation struct {
	Packages []Package

	fd   *os.File
	once sync.Once
}

func NewDerivation(path string) *Derivation {
	dpath := filepath.Join(path, "derivation.nix")
	fd, _ := os.OpenFile(dpath, os.O_RDWR|os.O_CREATE, 0755)
	return &Derivation{fd: fd}
}

func (d *Derivation) Execute(args []string, opts ...executer.Option) (int, error) {
	d.once.Do(d.generate)
	cmd := strings.Join(args, " ")
	exe := exec.Command("nix-shell", d.Path(), "--command", cmd)

	for _, opt := range opts {
		opt(exe)
	}

	err := exe.Run()
	code := exe.ProcessState.ExitCode()
	return code, err
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

	d.fd.Truncate(0)
	d.fd.Write(buf.Bytes())
	d.fd.Sync()
}

func (d *Derivation) Install() error {
	_, err := d.Execute([]string{"true"}, executer.WithStdout(os.Stdout))
	return err
}

func (d *Derivation) Path() string {
	return d.fd.Name()
}
