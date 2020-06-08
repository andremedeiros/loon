package config

import (
	"regexp"
	"strings"
)

var tldRegex = regexp.MustCompile(`^(\.[a-z][a-z0-9-]+)+$`)

type devTLDField string

func (f devTLDField) Validate() error {
	if tldRegex.Match([]byte(f)) {
		return nil
	}

	return ErrInvalidDevTLD
}

func (f devTLDField) Normalize() string {
	return strings.ToLower(string(f))
}
