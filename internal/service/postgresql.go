package service

type Postgres struct{}

func (p *Postgres) String() string {
	return "Postgres"
}

func (p *Postgres) Environ() []string {
	return []string{}
}

func (p *Postgres) Start() []string {
	return nil
}

func (p *Postgres) Stop() error {
	return nil
}
