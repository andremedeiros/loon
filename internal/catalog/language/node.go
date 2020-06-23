package language

import "github.com/andremedeiros/loon/internal/executer"

type Node struct{}

func (n *Node) String() string {
	return "Node"
}

func (n *Node) Environ() []string {
	return nil
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
