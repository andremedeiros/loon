package language

type Ruby struct{}

func (r *Ruby) String() string {
	return "Ruby"
}

func (r *Ruby) Environ() []string {
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

func (r *Ruby) Initialize(exe Executer, _ string) error {
	return exe.Execute([]string{
		"ruby",
		"--version",
	})
}
