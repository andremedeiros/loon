package catalog

type Postgres struct{}

func (p *Postgres) String() string {
	return "Postgres"
}

func (p *Postgres) Environ(ipaddr, vdpath string) []string {
	return []string{}
}

func (p *Postgres) Start(ipaddr, vdpath string) []string {
	return nil
}

func (p *Postgres) Stop(ipaddr, vdpath string) error {
	return nil
}
