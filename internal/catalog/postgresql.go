package catalog

type Postgres struct{}

func (p *Postgres) String() string {
	return "Postgres"
}

func (p *Postgres) Initialize(_, _ string) []string {
	return nil
}

func (p *Postgres) Versions() map[string]Entry {
	return map[string]Entry{
		"default": EntryFor("postgresql", "12.3"),
		"latest":  EntryFor("postgresql", "12.3"),

		"12.3": EntryFor("postgresql", "12.3"),
	}
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
