package service

type Memcached struct{}

func (m *Memcached) String() string {
	return "Memcached"
}

func (m *Memcached) Environ() []string {
	return []string{"MEMCACHED_URL=localhost:11211"}
}

func (m *Memcached) Start() []string {
	return []string{"memcached"}
}

func (m *Memcached) Stop() error {
	return nil
}
