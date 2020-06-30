package language

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/executor"
)

type Ruby struct{}

func (r *Ruby) String() string {
	return "Ruby"
}

func (r *Ruby) gemHome(vdpath string) string {
	return filepath.Join(vdpath, "data", "gem")
}

func (r *Ruby) Environ(vdpath string) []string {
	gemHome := r.gemHome(vdpath)
	return []string{
		fmt.Sprintf("GEM_HOME=%s", gemHome),
		fmt.Sprintf("GEM_PATH=%s", gemHome),
	}
}

func (r *Ruby) BinPaths(vdpath string) []string {
	gemBin := filepath.Join(r.gemHome(vdpath), "bin")
	return []string{gemBin}
}

func (r *Ruby) Versions() map[string][]string {
	return map[string][]string{
		"default": {"ruby", "2.7.1"},
		"latest":  {"ruby", "2.7.1"},

		"2.7.1": {"ruby", "2.7.1"},
		"2.6.6": {"ruby", "2.6.6"},
	}
}

func (r *Ruby) Initialize(exe executor.Executor, vdpath string, opts ...executor.Option) error {
	bundlerBinPath := filepath.Join(r.gemHome(vdpath), "bin", "bundle")
	if _, err := os.Stat(bundlerBinPath); os.IsNotExist(err) {
		return exe.Execute([]string{"gem", "install", "bundler", "--no-document"})
	}
	return nil
}
