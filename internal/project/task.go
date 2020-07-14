package project

type Tasks []Task

func (tt Tasks) Len() int {
	return len(tt)
}

func (tt Tasks) Swap(i, j int) {
	tt[i], tt[j] = tt[j], tt[i]
}

func (tt Tasks) Less(i, j int) bool {
	return tt[i].Name < tt[j].Name
}

type Task struct {
	Name        string
	Description string `yaml:"desc"`
	Command     string `yaml:"cmd"`
}
