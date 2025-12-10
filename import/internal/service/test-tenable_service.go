package service

import (
	"github.com/soner3/tenable-import/import/internal/config"
	"github.com/soner3/tenable-import/import/internal/repository"
)

// TestTenableServiceImpl implementiert den TenableService interface
type TestTenableServiceImpl struct {
	App           *config.AppConfig
	PiaRepository repository.PiaRepository
}

// NewTestTenableService erstellt eine neue Instanz von TestTenableServiceImpl
func NewTestTenableService(appConfig *config.AppConfig, piaRepo repository.PiaRepository) VulnerabilityService {
	return &TestTenableServiceImpl{
		App:           appConfig,
		PiaRepository: piaRepo,
	}
}
