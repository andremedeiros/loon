package executor

type Executor interface {
	Execute([]string, ...Option) (int, error)
}
