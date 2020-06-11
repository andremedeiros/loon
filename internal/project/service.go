package project

import (
	"fmt"

	"github.com/andremedeiros/loon/internal/provider"
	"github.com/andremedeiros/loon/internal/service"
)

type Service struct {
	Provider provider.Provider
	Service  service.Service
	Version  string
}

func (s *Service) String() string {
	return fmt.Sprintf("%s %s", s.Service.String(), s.Version)
}

func (s *Service) Start() error {
	cmd := s.Service.Start()
	return s.Provider.Start(cmd)
}

func (s *Service) Stop() {
}
