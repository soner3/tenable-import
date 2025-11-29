package service

import (
	"github.com/soner3/tenable-import/internal/config"
	"github.com/soner3/tenable-import/internal/repository"
)

// TenableServiceImpl implementiert den TenableService interface
type TenableServiceImpl struct {
	App           *config.AppConfig
	PiaRepository repository.PiaRepository
}

// NewTenableService erstellt eine neue Instanz von TenableServiceImpl
func NewTenableService(appConfig *config.AppConfig, piaRepo repository.PiaRepository) TenableService {
	return &TenableServiceImpl{
		App:           appConfig,
		PiaRepository: piaRepo,
	}
}

func (s *TenableServiceImpl) CreateAssets() error {
	return nil
}

func (s *TenableServiceImpl) DeleteAssets() {
}

func (s *TenableServiceImpl) GetAllAssets() {
}

func (s *TenableServiceImpl) CompareDBAndTenableAssets() {
}
