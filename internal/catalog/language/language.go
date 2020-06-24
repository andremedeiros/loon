package language

import "github.com/andremedeiros/loon/internal/executer"

type Language interface {
	String() string
	Environ(string) []string
	BinPaths(string) []string
	Initialize(executer.Executer, string, ...executer.Option) error
}

var Languages = map[string]Language{
	"crystal": &Crystal{},
	"golang":  &Golang{},
	"node":    &Node{},
	"ruby":    &Ruby{},
}
