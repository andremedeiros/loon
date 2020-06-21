package config

import (
	"testing"

	"gopkg.in/yaml.v2"
)

func TestConfigValidations(t *testing.T) {
	var cfgtests = []struct {
		in    string
		valid bool
	}{
		{`dev_tld: .something`, true},
		{`dev_tld: .some.sub.domains`, true},
		{`dev_tld: .f`, false},
		{`dev_tld: invalid`, false},
		{`source_tree: "$HOME/{name}"`, true},
		{`source_tree: "$HOME/src/{owner}/{name}"`, true},
		{`source_tree: "$HOME/src/{host}/{owner}/{name}"`, true},
		{`source_tree: "$HOME/src/{name}/{nonExistant}"`, true},
		{`source_tree: "$HOME/src"`, false},
	}
	for _, tt := range cfgtests {
		t.Run(tt.in, func(t *testing.T) {
			cfg := newDefaultConfig()
			yaml.Unmarshal([]byte(tt.in), cfg)
			err := cfg.Validate()
			if tt.valid && err != nil {
				t.Errorf("got %q", err)
			} else if !tt.valid && err == nil {
				t.Errorf("expected error but didn't get one")
			}
		})
	}
}
