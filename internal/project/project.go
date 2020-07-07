package project

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/andremedeiros/loon/internal/executor"
	"github.com/andremedeiros/loon/internal/nix"

	"gopkg.in/yaml.v3"
)

type Project struct {
	Name         string
	URL          string
	Provider     string
	Dependencies []Dependency
	Tasks        []Task
	Environment  map[string]string
	Path         string
	ModTime      time.Time
	IP           net.IP

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

func (p *Project) Execute(args []string, opts ...executor.Option) error {
	return p.derivation.Execute(args, opts...)
}

func (p *Project) NeedsUpdate() bool {
	return p.derivation.NeedsUpdate(p.ModTime)
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

func (p *Project) ensurePaths() {
	paths := []string{
		p.VariableDataPath(),
		p.VariableDataPath("pids"),
		p.VariableDataPath("sockets"),
		p.VariableDataPath("data"),
	}
	for _, p := range paths {
		os.MkdirAll(p, 0755)
	}
	gitignorePath := filepath.Join(p.VDPath(), ".gitignore")
	if _, err := os.Stat(gitignorePath); os.IsNotExist(err) {
		ioutil.WriteFile(gitignorePath, []byte{'*'}, 0600)
	}
}

func (p *Project) HostExecutor() executor.Executor {
	return p
}

func (p *Project) DerivationExecutor() executor.Executor {
	return p.derivation
}

func (p *Project) VariableDataPath(parts ...string) string {
	parts = append([]string{p.Path, ".loon"}, parts...)
	return filepath.Join(parts...)
}

func (p *Project) VDPath() string {
	return filepath.Join(p.Path, ".loon")
}

func (p *Project) UnmarshalYAML(unmarshal func(interface{}) error) error {
	pd := struct {
		Name        string
		URL         string
		Environment map[string]string
		Deps        []interface{}
		Tasks       map[string]map[string]string
	}{}

	if err := unmarshal(&pd); err != nil {
		return err
	}

	p.Name = pd.Name
	p.URL = pd.URL
	p.Environment = pd.Environment

	for _, dep := range pd.Deps {
		name := ""
		version := "default"
		switch dep.(type) {
		case string:
			name = dep.(string)
		case map[string]interface{}:
			for n, v := range dep.(map[string]interface{}) {
				name = n
				switch v.(type) {
				case string:
					version = v.(string)
				case float64:
					version = fmt.Sprintf("%g", v)
				}
			}
		default:
			return fmt.Errorf("specify dependencies as `dep` or `dep: version`")
		}
		pkg, err := nix.PackageFor(name, version)
		if err != nil {
			return err
		}
		p.derivation.Packages = append(p.derivation.Packages, pkg)
		d := Dependency{name, version, pkg}
		p.Dependencies = append(p.Dependencies, d)
	}
	for tn, opts := range pd.Tasks {
		t := Task{Name: tn}
		command, ok := opts["command"]
		if !ok {
			return fmt.Errorf("need command for task %s", tn)
		}
		t.Command = command
		t.Description = opts["description"]
		p.Tasks = append(p.Tasks, t)
	}

	return nil
}

func IsNotFound(err error) bool {
	return (err == ErrProjectPayloadNotFound)
}
