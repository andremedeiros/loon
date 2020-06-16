package nix

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

type Derivation struct {
	Packages []Package
	vdpath   string
}

func NewDerivation() *Derivation {
	return &Derivation{}
}

func (d *Derivation) Execute(args []string) error {
	cmd := strings.Join(args, " ")
	fmt.Println(cmd)
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
{{range $package := $.Packages}}
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

	fd, _ := os.OpenFile(d.Path(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer fd.Close()
	fd.Write(buf.Bytes())

	exe := exec.Command("nix-shell", d.Path(), "--command", "true")
	exe.Stdout = os.Stdout
	exe.Stderr = os.Stderr
	return exe.Run()
}

func (d *Derivation) Path() string {
	return filepath.Join("/tmp", "default.nix")
}
