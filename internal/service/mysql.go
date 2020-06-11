package service

type Mysql struct{}

func (m *Mysql) String() string {
	return "MySQL"
}

func (m *Mysql) Environ() []string {
	return []string{}
}

func (m *Mysql) Start() []string {
	return nil
}

func (m *Mysql) Stop() error {
	return nil
}
