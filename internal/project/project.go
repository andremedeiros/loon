package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/andremedeiros/loon/catalog"
	"github.com/andremedeiros/loon/internal/provider"
	"github.com/andremedeiros/loon/internal/service"
)

type Project struct {
	Name        string            `yaml:"name"`
	URL         string            `yaml:"url"`
	Provider    string            `yaml:"provider"`
	Services    []Service         `yaml:"services"`
	Environment map[string]string `yaml:"environment"`
}

func FindInTree() (*Project, error) {
	cwd, _ := os.Getwd()
	for cwd != filepath.Dir(cwd) {
		path := filepath.Join(cwd, "loon.yml")
		if _, err := os.Stat(path); os.IsNotExist(err) {
			cwd = filepath.Dir(cwd)
			continue
		}
		return NewFromPath(path)
	}
	return nil, ErrProjectPayloadNotFound
}

func NewFromPath(path string) (*Project, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewFromPayload(body)
}

func NewFromPayload(b []byte) (*Project, error) {
	project := &Project{}
	if err := yaml.Unmarshal(b, project); err != nil {
		return nil, err
	}
	return project, nil
}

func (p *Project) Environ() []string {
	environ := os.Environ()
	for _, s := range p.Services {
		environ = append(environ, s.Service.Environ()...)
	}
	for k, v := range p.Environment {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}
	return environ
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
		srvc, ok := service.Handlers[serviceName]
		if !ok {
			return fmt.Errorf("service not supported: %s", srvc)
		}

		ce, err := catalog.EntryFor("nix", serviceName, opts["version"])
		if err != nil {
			return err
		}

		prov := provider.Handlers["nix"]
		prov.Add(ce)

		s := Service{
			Provider: prov,
			Service:  srvc,
			Version:  ce.Version,
		}
		p.Services = append(p.Services, s)
	}

	return nil
}
