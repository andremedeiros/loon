package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/andremedeiros/loon/internal/catalog"
	"github.com/andremedeiros/loon/internal/catalog/language"
	"github.com/andremedeiros/loon/internal/catalog/service"
	"github.com/andremedeiros/loon/internal/nix"
)

type Project struct {
	Name        string              `yaml:"name"`
	URL         string              `yaml:"url"`
	Provider    string              `yaml:"provider"`
	Services    []service.Service   `yaml:"services"`
	Languages   []language.Language `yaml:"languages"`
	Environment map[string]string   `yaml:"environment"`
	Path        string

	derivation *nix.Derivation
}

func (p *Project) IPAddr() string {
	// TODO(andremedeiros): implement the thing
	return "127.0.0.1"
}

func (p *Project) Execute(args []string) error {
	return p.derivation.Execute(args)
}

func (p *Project) EnsureDependencies() error {
	return p.derivation.Install()
}

func FindInTree() (*Project, error) {
	cwd, _ := os.Getwd()
	for cwd != filepath.Dir(cwd) {
		path := filepath.Join(cwd, "loon.yml")
		if _, err := os.Stat(path); os.IsNotExist(err) {
			cwd = filepath.Dir(cwd)
			continue
		}

		p, err := fromPath(path)
		if err != nil {
			return nil, err
		}
		p.ensurePaths()
		return p, nil
	}
	return nil, ErrProjectPayloadNotFound
}

func fromPath(path string) (*Project, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	p, err := fromPayload(body)
	if err != nil {
		return nil, err
	}
	p.Path = filepath.Dir(path)
	return p, nil
}

func fromPayload(b []byte) (*Project, error) {
	project := &Project{derivation: nix.NewDerivation()}
	if err := yaml.Unmarshal(b, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (p *Project) Environ() []string {
	environ := os.Environ()
	for _, s := range p.Services {
		environ = append(environ, s.Environ(p.IPAddr(), p.VDPath())...)
	}
	for k, v := range p.Environment {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}
	return environ
}

func (p *Project) ensurePaths() {
	paths := []string{
		p.VDPath(),
		filepath.Join(p.VDPath(), "pids"),
		filepath.Join(p.VDPath(), "sockets"),
	}

	for _, svc := range p.Services {
		svcPath := filepath.Join(p.VDPath(), "data", svc.Identifier())
		paths = append(paths, svcPath)
	}

	for _, p := range paths {
		os.MkdirAll(p, 0755)
	}
}

func (p *Project) VDPath() string {
	return filepath.Join(p.Path, ".loon")
}

func (p *Project) UnmarshalYAML(unmarshal func(interface{}) error) error {
	projectData := struct {
		Name        string
		URL         string
		Environment map[string]string
	}{}

	if err := unmarshal(&projectData); err != nil {
		return err
	}

	p.Name = projectData.Name
	p.URL = projectData.URL
	p.Environment = projectData.Environment

	serviceData := struct {
		Services map[string]map[string]string
	}{}

	if err := unmarshal(&serviceData); err != nil {
		return err
	}

	for serviceName, opts := range serviceData.Services {
		version := "default"
		if _, ok := opts["version"]; ok {
			version = opts["version"]
		}

		svc, ok := service.Services[serviceName]
		if !ok {
			return fmt.Errorf("service not supported: %s", serviceName)
		}

		pkgs := catalog.Packages(svc.(catalog.Installable), version)
		if len(pkgs) == 0 {
			return fmt.Errorf("service not present in catalog: %s %s", serviceName, version)
		}

		for _, nixpkg := range pkgs {
			p.derivation.Packages = append(p.derivation.Packages, nixpkg)
		}

		p.Services = append(p.Services, svc)
	}

	languageData := struct {
		Languages map[string]map[string]string
	}{}

	if err := unmarshal(&languageData); err != nil {
		return err
	}

	for languageName, opts := range languageData.Languages {
		version := "default"
		if _, ok := opts["version"]; ok {
			version = opts["version"]
		}

		lang, ok := language.Languages[languageName]
		if !ok {
			return fmt.Errorf("language not supported: %s", languageName)
		}

		pkgs := catalog.Packages(lang.(catalog.Installable), version)
		if len(pkgs) == 0 {
			return fmt.Errorf("language not present in catalog: %s %s", languageName, version)
		}

		for _, nixpkg := range pkgs {
			p.derivation.Packages = append(p.derivation.Packages, nixpkg)
		}

		p.Languages = append(p.Languages, lang)
	}

	return nil
}
