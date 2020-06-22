package executer

type Executer interface {
	Execute([]string, ...Option) (int, error)
}
