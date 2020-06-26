package project

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v2"

	"github.com/andremedeiros/loon/internal/catalog"
	"github.com/andremedeiros/loon/internal/catalog/language"
	"github.com/andremedeiros/loon/internal/catalog/service"
	"github.com/andremedeiros/loon/internal/executer"
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
	ModTime     time.Time
	IP          net.IP

	derivation *nix.Derivation
}

func (p *Project) Task(name string) (Task, error) {
	for _, task := range p.Tasks {
		if task.Name == name {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("task not found: %s", name)
}

func ipFromPath(path string) net.IP {
	crc32q := crc32.MakeTable(0xD5828281)
	crc := crc32.Checksum([]byte(path), crc32q) % (255 * 255)
	part3 := crc / 255
	part4 := crc - (part3 * 255)
	addr := fmt.Sprintf("127.0.%d.%d", part3, part4)
	return net.ParseIP(addr)
}

func (p *Project) Execute(args []string, opts ...executer.Option) (int, error) {
	opts = append(opts, executer.WithEnviron(p.Environ()))
	return p.derivation.Execute(args, opts...)
}

func (p *Project) NeedsUpdate() bool {
	return p.derivation.NeedsUpdate(p.ModTime)
}

func (p *Project) EnsureDependencies() error {
	if !p.derivation.NeedsUpdate(p.ModTime) {
		return nil
	}
	return p.derivation.Install()
}

func (p *Project) NeedsNetworking() bool {
	ifis, err := net.Interfaces()
	if err != nil {
		return true
	}
	for _, ifi := range ifis {
		addrs, err := ifi.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ipnet.IP.Equal(p.IP) {
					return false
				}
			}
		}
	}
	return true
}

// TODO(andremedeiros): extract this into an OS dependent implementation
func (p *Project) EnsureNetworking(opts ...executer.Option) error {
	_, err := p.derivation.Execute([]string{
		"sudo",
		"ifconfig",
		"lo0",
		"alias",
		p.IP.String(),
		"255.255.255.0",
	}, opts...)
	return err
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
	p := &Project{}
	p.Path = filepath.Dir(path)
	p.IP = ipFromPath(path)
	p.derivation = nix.NewDerivation(p.VDPath())
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	p.ModTime = fi.ModTime()
	b, err := ioutil.ReadFile(path)
	if err := yaml.Unmarshal(b, p); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Project) Environ() []string {
	environ := os.Environ()
	paths := []string{}
	for _, s := range p.Services {
		environ = append(environ, s.Environ(p.IP, p.VDPath())...)
	}
	for _, l := range p.Languages {
		environ = append(environ, l.Environ(p.VDPath())...)
		paths = append(paths, l.BinPaths(p.VDPath())...)
	}
	for k, v := range p.Environment {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}
	path := fmt.Sprintf("PATH=%s:%s", strings.Join(paths, ":"), os.Getenv("PATH"))
	environ = append(environ, path)
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
