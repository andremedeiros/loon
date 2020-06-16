package project

import (
	"fmt"

	"github.com/andremedeiros/loon/internal/catalog"
)

type Service struct {
	Service catalog.Service
	Version string
}

func (s *Service) String() string {
	return fmt.Sprintf("%s %s", s.Service.String(), s.Version)
}

func (s *Service) Start() error {
	return nil
}

func (s *Service) Stop() error {
	return nil
}
