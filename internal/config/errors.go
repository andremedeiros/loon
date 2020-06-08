package config

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidSyntax             = errors.New("configuration has invalid syntax")
	ErrInvalidDevTLD             = errors.New("invalid dev TLD")
	ErrProviderNotSupported      = errors.New("provider not supported")
	ErrRepositoryNotInSourceTree = errors.New("repository placeholder not in source tree")
)

type ErrInvalidValueForConfig struct {
	field string
	err   error
}

func (e ErrInvalidValueForConfig) Error() string {
	return fmt.Sprintf("invalid configuration value for %+q", e.field)
}
