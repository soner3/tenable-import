package service

import (
	"slices"

	"github.com/soner3/tenable-import/internal/config"
	"github.com/soner3/tenable-import/internal/model"
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

func (s *TenableServiceImpl) CompareDBAndTenableAssets(dbAssets *[]model.Asset, tenableAssets *[]model.TenableExportResponseAsset) (foundAssets *[]model.Asset, missingAssets *[]model.Asset) {
	foundAssets = &[]model.Asset{}
	missingAssets = &[]model.Asset{}

	for _, dbAsset := range *dbAssets {
		found := false
		for _, tenableAsset := range *tenableAssets {
			found = slices.Contains(tenableAsset.Network.IPv4s, dbAsset.IPv4s[0])
			if !found {
				for _, ni := range tenableAsset.Network.NetworkInterfaces {
					found = slices.Contains(ni.IPv4s, dbAsset.IPv4s[0])
					if found {
						break
					}
				}
			}

			if found {
				*foundAssets = append(*foundAssets, dbAsset)
				break
			}
		}

		if !found {
			*missingAssets = append(*missingAssets, dbAsset)
		}
	}

	return foundAssets, missingAssets
}
