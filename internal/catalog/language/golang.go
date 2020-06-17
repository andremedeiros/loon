package language

type Golang struct{}

func (g *Golang) String() string {
	return "Golang"
}

func (g *Golang) Environ() []string {
	return nil
}

func (g *Golang) Versions() map[string][]string {
	return map[string][]string{
		"default": {"golang", "1.14.4"},
		"latest":  {"golang", "1.14.4"},

		"1.14.4":  {"golang", "1.14.4"},
		"1.13.12": {"golang", "1.13.12"},
	}
}

func (g *Golang) Initialize(exe Executer, _ string) error {
	return exe.Execute([]string{
		"go",
		"version",
	})
}
