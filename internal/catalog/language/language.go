package language

type Executer interface {
	Execute([]string) error
}

type Language interface {
	String() string
	Environ() []string
	Initialize(Executer, string) error
}

var Languages = map[string]Language{
	"ruby":   &Ruby{},
	"golang": &Golang{},
}
