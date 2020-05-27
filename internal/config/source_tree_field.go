package config

import (
	"fmt"
	"os"
	"strings"
)

type sourceTreeField string

func (f sourceTreeField) Validate() error {
	if f.Resolve("", "", "A") == f.Resolve("", "", "B") {
		return ErrRepositoryNotInSourceTree
	}
	return nil
}

func (f sourceTreeField) Resolve(host, owner, name string) string {
	rs := []string{
		"{host}", host,
		"{owner}", owner,
		"{name}", name,
	}
	for _, v := range os.Environ() {
		parts := strings.SplitN(v, "=", 2)
		rs = append(rs, fmt.Sprintf("$%s", parts[0]), parts[1])
	}
	return strings.NewReplacer(rs...).Replace(string(f))
}
