package project

type Task struct {
	Name        string
	Description string `yaml:"desc"`
	Command     string `yaml:"cmd"`
}
