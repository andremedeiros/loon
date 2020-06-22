package nix

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"sync"
	"text/template"
)

type Derivation struct {
	Packages []Package

	tmpfile *os.File
	once    sync.Once
}

func NewDerivation() *Derivation {
	tmpfile, _ := ioutil.TempFile("", "derivation.nix")
	return &Derivation{tmpfile: tmpfile}
}

func (d *Derivation) Execute(args []string) error {
	d.once.Do(d.generate)
	cmd := strings.Join(args, " ")
	exe := exec.Command("nix-shell", d.Path(), "--command", cmd)

	if true {
		exe.Stdout = os.Stdout
		exe.Stderr = os.Stderr
	}

	return exe.Run()
}

func (d *Derivation) generate() {
	tmpl := `
{ pkgs ? import <nixpkgs> { } }:
let
inherit (pkgs) fetchurl mkShell;
{{ range $package := $.Packages }}
	{{ if eq $package.Inherit "" }}
		{{ $package.Name }} = pkgs.{{ $package.Name }}.overrideAttrs (attrs: {
			version = "{{ $package.Version }}";
			src = fetchurl {
				url = "{{ $package.URL }}";
				sha256 = "{{ $package.SHA256 }}";
			};
		});
	{{ else }}
		{{ $package.Name }} = pkgs.{{ $package.Inherit }};
	{{ end }}
{{ end }}

in mkShell {
	buildInputs = [{{ range $package := .Packages }}
		{{ $package.Name }}
	{{ end }}];
}`

	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("nix").Parse(tmpl))
	t.Execute(buf, d)

	d.tmpfile.Truncate(0)
	d.tmpfile.Write(buf.Bytes())
	d.tmpfile.Sync()
}

func (d *Derivation) Install() error {
	return d.Execute([]string{"true"})
}

func (d *Derivation) Path() string {
	return d.tmpfile.Name()
}
