package catalog

type Memcached struct{}

func (m *Memcached) String() string {
	return "Memcached"
}

func (m *Memcached) Environ(ipaddr, vdpath string) []string {
	return []string{"MEMCACHED_URL=localhost:11211"}
}

func (m *Memcached) Start(ipaddr, vdpath string) []string {
	return []string{"memcached"}
}

func (m *Memcached) Stop(ipaddr, vdpath string) error {
	return nil
}
