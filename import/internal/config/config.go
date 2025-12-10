package config

import "github.com/soner3/tenable-import/lib/config"

// AppConfig erweitert die Basis-Konfiguration um anwendungsspezifische Felder
type AppConfig struct {
	config.Config
	DSN           string
	TenableAPIKey string
}

// NewAppConfig erstellt eine neue AppConfig mit Standardwerten
func NewAppConfig() *AppConfig {
	return &AppConfig{
		Config: config.NewConfig(),
	}
}
