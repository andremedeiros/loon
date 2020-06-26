package language

import "github.com/andremedeiros/loon/internal/executor"

type Language interface {
	String() string
	Environ(string) []string
	BinPaths(string) []string
	Initialize(executor.Executor, string, ...executor.Option) error
}

var Languages = map[string]Language{
	"crystal": &Crystal{},
	"golang":  &Golang{},
	"node":    &Node{},
	"ruby":    &Ruby{},
}
