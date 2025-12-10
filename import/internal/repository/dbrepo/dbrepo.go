package dbrepo

import (
	"github.com/soner3/tenable-import/import/internal/config"
	"github.com/soner3/tenable-import/import/internal/repository"
)

// piaRepositoryImpl implementiert das PiaRepository interface
type piaRepositoryImpl struct {
	App *config.AppConfig
}

// NewPiaRepository erstellt eine neue Instanz von piaRepositoryImpl
func NewPiaRepository(appConfig *config.AppConfig) repository.PiaRepository {
	return &piaRepositoryImpl{
		App: appConfig,
	}
}

// testPiaRepository implementiert das PiaRepository interface für Testzwecke
type testPiaRepository struct {
	App *config.AppConfig
}

// NewTestPiaRepository erstellt eine neue Instanz von testPiaRepository
func NewTestPiaRepository(appConfig *config.AppConfig) repository.PiaRepository {
	return &testPiaRepository{
		App: appConfig,
	}
}
