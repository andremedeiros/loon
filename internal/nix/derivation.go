package nix

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/andremedeiros/loon/internal/catalog"
)

type Derivation struct {
	Entries []catalog.Entry
	vdpath  string
}

func NewDerivation(vdpath string) *Derivation {
	return &Derivation{vdpath: vdpath}
}

func (d *Derivation) Execute(args []string) error {
	cmd := strings.Join(args, " ")
	exe := exec.Command("nix-shell", d.Path(), "--command", cmd)
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	return exe.Run()
}

func (d *Derivation) Install() error {
	tmpl := `
{ pkgs ? import <nixpkgs> { } }:
let
inherit (pkgs) fetchurl mkShell;
{{range $entry := .Entries}}
{range $package := $entry.Packages}}
	{{$package.Name}} = pkgs.{{$package.Name}}.overrideAttrs (attrs: {
		version = "{{$package.Version}}";
		src = fetchurl {
			url = "{{$package.URL}}";
			sha256 = "{{$package.SHA256}}";
		};
	});
{{end}}

in mkShell {
	buildInputs = [{{range $package := .Packages}}
		{{$package.Name}}
	{{end}}];
}`

	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("nix").Parse(tmpl))
	t.Execute(buf, d)

	fd, _ := os.OpenFile(d.Path(), os.O_RDWR|os.O_CREATE, 0755)
	defer fd.Close()
	fd.Write(buf.Bytes())

	exe := exec.Command("nix-shell", d.Path(), "--command", "true")
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	return exe.Run()
}

func (d *Derivation) Path() string {
	return filepath.Join(d.vdpath, "default.nix")
}
