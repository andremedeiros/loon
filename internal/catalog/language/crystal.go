package language

import "github.com/andremedeiros/loon/internal/executor"

type Crystal struct{}

func (c *Crystal) String() string {
	return "Crystal"
}

func (c *Crystal) Environ(_ string) []string {
	return nil
}

func (c *Crystal) BinPaths(_ string) []string {
	return nil
}

func (c *Crystal) Versions() map[string][]string {
	return map[string][]string{
		"default": {"crystal", "0.35.1"},
		"latest":  {"crystal", "0.35.1"},

		"0.35.1": {"crystal", "0.35.1"},
	}
}

func (c *Crystal) Initialize(exe executor.Executor, _ string, opts ...executor.Option) error {
	return nil
}
