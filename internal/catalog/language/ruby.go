package language

import "github.com/andremedeiros/loon/internal/executor"

type Ruby struct{}

func (r *Ruby) String() string {
	return "Ruby"
}

func (r *Ruby) Environ(_ string) []string {
	return nil
}

func (r *Ruby) BinPaths(_ string) []string {
	return nil
}

func (r *Ruby) Versions() map[string][]string {
	return map[string][]string{
		"default": {"ruby", "2.7.1"},
		"latest":  {"ruby", "2.7.1"},

		"2.7.1": {"ruby", "2.7.1"},
		"2.6.6": {"ruby", "2.6.6"},
	}
}

func (r *Ruby) Initialize(exe executor.Executor, _ string, opts ...executor.Option) error {
	_, err := exe.Execute([]string{"gem", "install", "bundler", "--no-document"})
	return err
}
