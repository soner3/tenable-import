package repository

import "github.com/soner3/tenable-import/model"

// PiaRepository definiert die Methoden für die Interaktion mit PIA-Assets
type PiaRepository interface {
	GetAllOpenAssets() ([]*model.Asset, error)
}
