package service

import (
	"github.com/soner3/tenable-import/import/internal/config"
	"github.com/soner3/tenable-import/import/internal/repository"
	"github.com/soner3/tenable-import/lib/helper"
	"github.com/soner3/tenable-import/lib/tenable"
)

// TenableServiceImpl implementiert den TenableService interface
type TenableServiceImpl struct {
	App           *config.AppConfig
	PiaRepository repository.PiaRepository
	Tenable       *tenable.TenableClient
}

// NewTenableService erstellt eine neue Instanz von TenableServiceImpl
func NewTenableService(appConfig *config.AppConfig, piaRepo repository.PiaRepository) VulnerabilityService {
	return &TenableServiceImpl{
		App:           appConfig,
		PiaRepository: piaRepo,
	}
}

// GetAllAssets ruft alle Assets von Tenable ab
func (s *TenableServiceImpl) GetAllAssets() ([]*tenable.Asset, error) {
	s.App.InfoLogger.Println("=== Tenable Asset-Import Service gestartet ===")
	s.App.DebugLogger.Println("Bereite Asset-Export vor...")

	filters := &tenable.AssetExportFilters{}

	s.App.DebugLogger.Println("Rufe Tenable API auf...")
	assets, err := s.Tenable.ExportAssetsV2(filters)
	if err != nil {
		s.App.ErrorLogger.Printf("Asset-Export fehlgeschlagen: %v", err)
		return nil, helper.WrapError(err, "Fehler beim Tenable Asset-Export")
	}

	if len(assets) == 0 {
		s.App.WarnLogger.Println("Keine Assets gefunden")
		return assets, nil
	}

	s.App.InfoLogger.Printf("=== Asset-Import erfolgreich: %d Assets geladen ===", len(assets))
	s.App.DebugLogger.Printf("Asset-Details: Erste Asset-ID: %s", assets[0].ID)

	return assets, nil
}
