package catalog

type Mysql struct{}

func (m *Mysql) String() string {
	return "MySQL"
}

func (m *Mysql) Environ(ipaddr, vdpath string) []string {
	return []string{}
}

func (m *Mysql) Start(ipaddr, vdpath string) []string {
	return nil
}

func (m *Mysql) Stop(ipaddr, vdpath string) error {
	return nil
}
