package task

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/andremedeiros/loon/internal/project"
)

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

func (*DotenvSetup) Environ(_ context.Context, p *project.Project) (Environ, BinPaths) {
	environ := []string{}
	for k, v := range p.Environment {
		environ = append(environ, fmt.Sprintf("%s=%s", k, v))
	}
	dotenv := filepath.Join(p.Path, ".env")
	if _, err := os.Stat(dotenv); os.IsNotExist(err) {
		return environ, nil
	}
	f, _ := os.Open(dotenv)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := strings.TrimSpace(scanner.Text())
		switch {
		case len(txt) == 0:
			continue
		case txt[0] == '#':
			continue
		default:
			environ = append(environ, txt)
		}
	}
	return environ, nil
}

func init() {
	RegisterTask("dotenv:setup", &DotenvSetup{})
	RunsAfter("command:up", "dotenv:setup")
	RunsAfter("command:exec", "dotenv:setup")
	RunsAfter("command:task", "dotenv:setup")
}
