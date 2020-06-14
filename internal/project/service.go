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

	IPAddress        string
	VariableDataPath string
}

func (s *Service) String() string {
	return fmt.Sprintf("%s %s", s.Service.String(), s.Version)
}

func (s *Service) Start() error {
	cmd := s.Service.Start(s.IPAddress, s.VariableDataPath)
	return s.Provider.Start(cmd)
}

func (s *Service) Stop() error {
	cmd := s.Service.Start(s.IPAddress, s.VariableDataPath)
	return s.Provider.Stop(cmd)
}
