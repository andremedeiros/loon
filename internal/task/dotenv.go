package task

import (
	"bufio"
	"context"
	"os"
	"path/filepath"
	"regexp"

	"github.com/andremedeiros/loon/internal/project"
)

var dotenvRe = regexp.MustCompile(`^[^#\w]*([A-Z0-9_]*)\=(\"(.*)\"|(.*))$`)

type DotenvSetup struct{}

func (*DotenvSetup) Header() string {
	return "Setting up environment..."
}

func (*DotenvSetup) Check(_ context.Context, p *project.Project) (bool, error) {
	return true, nil
}

func (*DotenvSetup) Resolve(_ context.Context, p *project.Project, _ SudoFunc) error {
	return nil
}

func (*DotenvSetup) Env(_ context.Context, p *project.Project) (Env, BinPaths) {
	environ := Env{"LOON_PROJECT_ROOT": p.Path}
	for k, v := range p.Environment {
		environ[k] = v
	}
	dotenv := filepath.Join(p.Path, ".env")
	if _, err := os.Stat(dotenv); os.IsNotExist(err) {
		return environ, nil
	}
	f, _ := os.Open(dotenv)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		match := dotenvRe.FindStringSubmatch(scanner.Text())
		if len(match) == 0 {
			continue
		}
		if match[3] != "" {
			environ[match[1]] = match[3]
		}
		if match[4] != "" {
			environ[match[1]] = match[4]
		}
	}
	return environ, nil
}

func init() {
	RegisterTask("dotenv:setup", &DotenvSetup{})
	RunsAfter("command:up:done", "dotenv:setup")
	RunsAfter("command:exec:done", "dotenv:setup")
	RunsAfter("command:task:done", "dotenv:setup")
}
