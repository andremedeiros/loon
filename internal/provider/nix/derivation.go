package nix

import (
	"bytes"
	"text/template"
)

type Derivation struct {
	Packages []Package
}

func (d *Derivation) Nix() string {
	tmpl := `
{ pkgs ? import <nixpkgs> { } }:
let
inherit (pkgs) fetchurl mkShell;
{{range $package := .Packages}}
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

	b := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("nix").Parse(tmpl))
	t.Execute(b, d)

	return b.String()
}
