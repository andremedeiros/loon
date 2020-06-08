package project

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Project struct {
	Name        string            `yaml:"name"`
	URL         string            `yaml:"url"`
	Provider    string            `yaml:"provider"`
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
	for k, v := range p.Environment {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}
	return environ
}
