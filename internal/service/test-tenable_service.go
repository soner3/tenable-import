package service

import (
	"github.com/soner3/tenable-import/internal/config"
	"github.com/soner3/tenable-import/internal/model"
	"github.com/soner3/tenable-import/internal/repository"
)

// TestTenableServiceImpl implementiert den TenableService interface
type TestTenableServiceImpl struct {
	App           *config.AppConfig
	PiaRepository repository.PiaRepository
}

// NewTestTenableService erstellt eine neue Instanz von TestTenableServiceImpl
func NewTestTenableService(appConfig *config.AppConfig, piaRepo repository.PiaRepository) TenableService {
	return &TestTenableServiceImpl{
		App:           appConfig,
		PiaRepository: piaRepo,
	}
}

func (s *TestTenableServiceImpl) CreateAssets() error {
	return nil
}

func (s *TestTenableServiceImpl) DeleteAssets() {
}

func (s *TestTenableServiceImpl) GetAllAssets() *[]model.TenableExportResponseAsset {
	return &[]model.TenableExportResponseAsset{
		{
			ID: "tenable-uuid-1111",
			Network: struct {
				IPv4s             []string `json:"ipv4s,omitempty"`
				FQDNs             []string `json:"fqdns,omitempty"`
				NetworkInterfaces []struct {
					IPv4s []string `json:"ipv4s,omitempty"`
					FQDNs []string `json:"fqdns,omitempty"`
				} `json:"network_interfaces,omitempty"`
			}{
				IPv4s: []string{"192.168.1.10"},
				FQDNs: []string{"server-existing.lan"},
			},
		},
		{
			ID: "tenable-uuid-3333",
			Network: struct {
				IPv4s             []string `json:"ipv4s,omitempty"`
				FQDNs             []string `json:"fqdns,omitempty"`
				NetworkInterfaces []struct {
					IPv4s []string `json:"ipv4s,omitempty"`
					FQDNs []string `json:"fqdns,omitempty"`
				} `json:"network_interfaces,omitempty"`
			}{
				NetworkInterfaces: []struct {
					IPv4s []string `json:"ipv4s,omitempty"`
					FQDNs []string `json:"fqdns,omitempty"`
				}{
					{
						IPv4s: []string{"192.168.1.30"},
						FQDNs: []string{"old-server.lan"},
					},
				},
			},
		},
	}
}

func (s *TestTenableServiceImpl) CompareDBAndTenableAssets() *[]model.Asset {
	return nil
}
