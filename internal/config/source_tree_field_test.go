package config

import (
	"os"
	"testing"
)

func TestSourceTreeFieldResolver(t *testing.T) {
	os.Setenv("HOME", "/home/user")
	var fieldtests = []struct {
		in  string
		out string
	}{
		{"$HOME/src", "/home/user/src"},
		{"$HOME/src/{host}/{owner}/{name}", "/home/user/src/gp/ao/repo"},
	}
	for _, tt := range fieldtests {
		t.Run(tt.in, func(t *testing.T) {
			f := sourceTreeField(tt.in)
			if path := f.Resolve("gp", "ao", "repo"); path != tt.out {
				t.Errorf("got %q, wanted %q", path, tt.out)
			}
		})
	}
}
