package project

import (
	"fmt"
	"hash/crc32"
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
	Tasks       []Task              `yaml:"tasks"`
	Environment map[string]string   `yaml:"environment"`
	Path        string

	derivation *nix.Derivation
}

func (p *Project) IPAddr() string {
	crc32q := crc32.MakeTable(0xD5828281)
	crc := crc32.Checksum([]byte(p.Path), crc32q) % (255 * 255)

	part3 := crc / 255
	part4 := crc - (part3 * 255)
	return fmt.Sprintf("127.0.%d.%d", part3, part4)
}

func (p *Project) Execute(args []string) error {
	return p.derivation.Execute(args, p.Environ())
}

func (p *Project) EnsureDependencies() error {
	return p.derivation.Install()
}

// TODO(andremedeiros): extract this into an OS dependent implementation
func (p *Project) EnsureNetworking() error {
	return p.derivation.Execute([]string{
		"sudo",
		"ifconfig",
		"lo0",
		"alias",
		p.IPAddr(),
		"255.255.255.0",
	}, nil)
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

	configData := struct {
		Languages map[string]map[string]string `yaml:"languages"`
		Services  map[string]map[string]string `yaml:"services"`
		Tasks     map[string]map[string]string `yaml:"tasks"`
	}{}

	if err := unmarshal(&configData); err != nil {
		return err
	}

	for serviceName, opts := range configData.Services {
		version := "default"
		if ver, ok := opts["version"]; ok {
			version = ver
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

	if err := unmarshal(&projectData); err != nil {
		return err
	}

	for languageName, opts := range configData.Languages {
		version := "default"
		if ver, ok := opts["version"]; ok {
			version = ver
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

	for taskName, opts := range configData.Tasks {
		task := Task{Name: taskName}
		command, ok := opts["command"]
		if !ok {
			return fmt.Errorf("need command for task %s", taskName)
		}
		task.Command = command
		task.Description = opts["description"]
		p.Tasks = append(p.Tasks, task)
	}

	return nil
}
