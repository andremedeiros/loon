package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/yaml.v3"
)

const (
	defaultDevTLD           = ".test"
	defaultSilenceDevDevTLD = false
	defaultSourceTree       = "$HOME/src/{host}/{owner}/{name}"
)

type Config struct {
	DevTLD           devTLDField     `yaml:"dev_tld"`
	SilenceDevDevTLD bool            `yaml:"silence_dev_dev_tld"`
	SourceTree       sourceTreeField `yaml:"source_tree"`
}

func newDefaultConfig() *Config {
	return &Config{
		DevTLD:           defaultDevTLD,
		SilenceDevDevTLD: defaultSilenceDevDevTLD,
		SourceTree:       defaultSourceTree,
	}
}

func xdgConfigHomePath(elem ...string) string {
	switch runtime.GOOS {
	case "darwin", "linux":
		return filepath.Join(os.Getenv("HOME"), ".config", filepath.Join(elem...))
	case "windows":
		return filepath.Join(os.Getenv("LOCALAPPDATA"), filepath.Join(elem...))
	}
	panic("operating system not supported")
}

// Read tries to load the configuration from disk, validates it, and returns it.
func Read() (*Config, error) {
	cfg := newDefaultConfig()

	path := xdgConfigHomePath("loon", "config.yml")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg, nil
	}

	body, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if yaml.Unmarshal(body, cfg) != nil {
		return nil, ErrInvalidSyntax
	}

	return cfg, cfg.Validate()
}

func (cfg *Config) Validate() error {
	if err := cfg.DevTLD.Validate(); err != nil {
		return ErrInvalidValueForConfig{"dev_tld", err}
	}
	if err := cfg.SourceTree.Validate(); err != nil {
		return ErrInvalidValueForConfig{"source_tree", err}
	}
	return nil
}
