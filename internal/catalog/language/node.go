package language

import (
	"fmt"
	"path/filepath"

	"github.com/andremedeiros/loon/internal/executer"
)

type Node struct{}

func (n *Node) String() string {
	return "Node"
}

func (n *Node) Environ(vdpath string) []string {
	npmPath := filepath.Join(vdpath, "data", "npm")
	return []string{
		fmt.Sprintf("NPM_CONFIG_PREFIX=%s", npmPath),
	}
}

func (n *Node) BinPaths(vdpath string) []string {
	return []string{
		filepath.Join(vdpath, "data", "npm", "bin"),
	}
}

func (n *Node) Versions() map[string][]string {
	return map[string][]string{
		"default": {"node", "12.18.1"},
		"latest":  {"node", "14.4.0"},

		"12.8.1": {"node", "12.18.1"},
		"14.4.0": {"node", "14.4.0"},
	}
}

func (n *Node) Initialize(exe executer.Executer, _ string, opts ...executer.Option) error {
	_, err := exe.Execute([]string{"npm", "install", "-g", "yarn"})
	return err
}
