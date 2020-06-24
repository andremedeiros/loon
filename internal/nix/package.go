package nix

import (
	"bytes"
	"fmt"
	"runtime"
	"text/template"
)

type PackageType int

const (
	Override PackageType = iota
	Inherit
	MakeDerivation
)

const (
	overrideTmpl = `
		{{ .Name }} = pkgs.{{ .Name }}.overrideAttrs(attrs: {
			version = "{{ .Version }}";
			src = fetchurl {
				url = "{{ .URL }}";
				sha256 = "{{ .SHA256 }}";
			};
		});`
	inheritTmpl = `
		{{ .Name }} = pkgs.{{ .Name }};`
	makeDerivationTmpl = `
    {{ .Name }} = stdenv.mkDerivation rec {
			pname = "{{ .Name }}";
			version = "{{ .Version }}";
			src = fetchurl {
				url = "{{ .URL }}";
				sha256 = "{{ .SHA256 }}";
			};
			buildCommand = ''
				{{ .BuildCommand }}
			'';
		};`
)

type Package struct {
	Name string
	Type PackageType

	Version string
	URL     string
	SHA256  string

	BuildCommand string
}

func NewPackage(name string, opts map[string]string) Package {
	platform := fmt.Sprintf("%s:%s", runtime.GOARCH, runtime.GOOS)
	p := Package{
		Name:    name,
		Version: opts["version"],
		URL:     opts["url"],
		SHA256:  opts["sha256"],
	}
	if sha, ok := opts[fmt.Sprintf("sha256:%s", platform)]; ok {
		p.SHA256 = sha
	}
	if url, ok := opts[fmt.Sprintf("url:%s", platform)]; ok {
		p.URL = url
	}
	switch opts["type"] {
	case "inherit":
		p.Type = Inherit
	case "derivation":
		p.Type = MakeDerivation
		p.BuildCommand = opts["buildCommand"]
	default:
		p.Type = Override
	}
	return p
}

func (p *Package) Nix() string {
	tmpl := ""
	switch p.Type {
	case Override:
		tmpl = overrideTmpl
	case Inherit:
		tmpl = inheritTmpl
	case MakeDerivation:
		tmpl = makeDerivationTmpl
	}
	buf := bytes.NewBuffer([]byte{})
	t := template.Must(template.New("nix").Parse(tmpl))
	t.Execute(buf, p)
	return buf.String()
}
