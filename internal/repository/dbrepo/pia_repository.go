package dbrepo

import "github.com/soner3/tenable-import/internal/model"

// GetAllAssets ruft alle PIA-Assets ab
func (r *piaRepositoryImpl) GetAllOpenAssets() ([]*model.Asset, error) {
	return []*model.Asset{}, nil
}

// GetAllClosedAssets ruft alle geschlossenen PIA-Assets ab
func (r *piaRepositoryImpl) GetAllClosedAssets() ([]*model.Asset, error) {
	return []*model.Asset{}, nil
}
