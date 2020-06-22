package language

import "github.com/andremedeiros/loon/internal/executer"

type Language interface {
	String() string
	Environ() []string
	Initialize(executer.Executer, string, ...executer.Option) error
}

var Languages = map[string]Language{
	"ruby":   &Ruby{},
	"golang": &Golang{},
}
